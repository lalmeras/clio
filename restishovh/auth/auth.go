package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/danielgtaylor/restish/cli"
	"github.com/spf13/cobra"
)

func loadCredentials() (*credentials, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configFile := path.Join(home, ".restish", "ovh-auth.json")
	f, err := os.OpenFile(configFile, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(f)
	credentials := &credentials{}
	err = dec.Decode(credentials)
	if err != nil {
		return nil, err
	}
	return credentials, nil
}

func Authenticate(cmd *cobra.Command, args []string) error {
	credentials, err := loadCredentials()
	if err != nil {
		return err
	}
	dec := json.NewDecoder(os.Stdin)
	input := &cli.Request{}
	if err = dec.Decode(input); err != nil {
		return err
	}
	result := signRequest(*input, *credentials)
	enc := json.NewEncoder(os.Stdout)
	if err = enc.Encode(result); err != nil {
		return err
	}
	return nil
}

func signRequest(request cli.Request, credentials credentials) cli.Request {
	result := cli.Request(request)
	headers := http.Header(request.Header)
	now := time.Now()
	headers.Set("X-Ovh-Application", credentials.ClientId)
	headers.Set("X-Ovh-Consumer", credentials.ConsumerKey)
	headers.Set("X-Ovh-Timestamp", fmt.Sprint(now.Unix()))
	headers.Set("X-Ovh-Signature", signPayload(credentials, result.Method, result.URI, result.Body, now))
	return result
}

func signPayload(credentials credentials, method string, url string, body string, timestamp time.Time) string {
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

type credentials struct {
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	ConsumerKey  string `json:"consumerKey,omitempty"`
}
