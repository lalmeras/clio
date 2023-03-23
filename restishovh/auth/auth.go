package auth

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/danielgtaylor/restish/cli"
)

func loadCredentials(configFile string) (*authConfig, error) {
	f, err := os.OpenFile(configFile, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(f)
	authConfig := &authConfig{}
	err = dec.Decode(authConfig)
	if err != nil {
		return nil, err
	}
	return authConfig, nil
}

func Authenticate(configFile string, account string) error {
	authConfig, err := loadCredentials(configFile)
	if err != nil {
		return err
	}
	credentials, ok := authConfig.Accounts[account]
	if !ok {
		return fmt.Errorf("Account %s not found in %s",
			account, configFile)
	}
	dec := json.NewDecoder(os.Stdin)
	input := &cli.Request{}
	if err = dec.Decode(input); err != nil {
		return err
	}
	result := signRequest(input, credentials)
	enc := json.NewEncoder(os.Stdout)
	if err = enc.Encode(result); err != nil {
		return err
	}
	return nil
}

type authHandler struct {
	c chan string
}

func (h authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := r.URL.Query().Get("error"); err != "" {
		rendered := "KO"
		w.Write([]byte(rendered))
		h.c <- "KO"
		return
	}

	h.c <- "OK"
	w.Write([]byte("OK"))
}

func Login(configFile string, account string) error {
	credentials, err := loadCredentials(configFile)
	payload := credentialRequest{}
	payload.AccessRules = append(payload.AccessRules, &accessRule{Method: "GET", Path: "/*"})
	payload.Redirection = "http://localhost:8080"
	c, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	b := ioutil.NopCloser(bytes.NewBuffer(c))
	client := &http.Client{}
	authUrl, _ := url.Parse("https://api.ovh.com/1.0/auth/credential")
	request := http.Request{
		URL:    authUrl,
		Method: "POST",
		Header: http.Header{
			"X-Ovh-Application": []string{credentials.Accounts[account].ClientId},
			"Content-Type":      []string{"application/json"},
		},
		Body: b,
	}
	resp, err := client.Do(&request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("/auth/credential request fails with error %d.", resp.StatusCode)
	}
	var credentialResponse credentialResponse
	json.NewDecoder(resp.Body).Decode(&credentialResponse)
	fmt.Printf("%s", credentialResponse.ConsumerKey)

	xdgOpen := exec.Command("xdg-open")
	xdgOpen.Args = append(xdgOpen.Args, credentialResponse.ValidationUrl)
	xdgOpen.Run()

	httpChan := make(chan string)
	authHandler := authHandler{c: httpChan}
	http.Handle("/", authHandler)
	go func() {
		// Run in a goroutine until the server is closed or we get an error.
		if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	var result = <-httpChan
	if result != "OK" {
		return fmt.Errorf("Authentication process failed.")
	}
	return nil
}

func signRequest(request *cli.Request, credentials *credentials) cli.Request {
	result := cli.Request(*request)
	headers := http.Header(request.Header)
	now := time.Now()
	headers.Set("X-Ovh-Application", credentials.ClientId)
	headers.Set("X-Ovh-Consumer", credentials.ConsumerKey)
	headers.Set("X-Ovh-Timestamp", fmt.Sprint(now.Unix()))
	headers.Set("X-Ovh-Signature", signPayload(credentials, result.Method, result.URI, result.Body, now))
	return result
}

func signPayload(credentials *credentials, method string, url string, body string, timestamp time.Time) string {
	trace, _ := os.OpenFile("/tmp/trace", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	defer trace.Close()
	message := strings.Join([]string{
		credentials.ClientSecret,
		credentials.ConsumerKey,
		strings.ToUpper(method),
		url,
		body,
		fmt.Sprint(timestamp.Unix()),
	}, "+")
	fmt.Fprint(trace, message)
	sum := sha1.Sum([]byte(message))
	return fmt.Sprintf("$1$%s", hex.EncodeToString(sum[:]))
}

type authConfig struct {
	Accounts map[string]*credentials `json:"accounts"`
}

type credentials struct {
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	ConsumerKey  string `json:"consumerKey,omitempty"`
}

type credentialRequest struct {
	AccessRules []*accessRule `json:"accessRules,omitempty"`
	Redirection string        `json:"redirection,omitempty"`
}

type accessRule struct {
	Method string `json:"method,omitempty"`
	Path   string `json:"path,omitempty"`
}

type credentialResponse struct {
	ConsumerKey   string          `json:"consumerKey,omitempty"`
	State         credentialState `json:"state,omitempty"`
	ValidationUrl string          `json:"validationUrl,omitempty"`
}

type credentialState string

const (
	expired           credentialState = "expired"
	pendingValidation                 = "pendingValidation"
	refused                           = "refused"
	validated                         = "validated"
)
