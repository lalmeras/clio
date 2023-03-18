package main

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

func Execute() {
	var genCmd = &cobra.Command{
		Use:   "restish-ovh",
		Short: "Restish external tool for OVH authentication",
		RunE:  Authenticate,
		Args:  cobra.NoArgs,
	}
	if err := genCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

func LoadCredentials() (*Credentials, error) {
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
	credentials := &Credentials{}
	err = dec.Decode(credentials)
	if err != nil {
		return nil, err
	}
	return credentials, nil
}

func Authenticate(cmd *cobra.Command, args []string) error {
	credentials, err := LoadCredentials()
	if err != nil {
		return err
	}
	dec := json.NewDecoder(os.Stdin)
	input := &cli.Request{}
	if err = dec.Decode(input); err != nil {
		return err
	}
	result := SignRequest(*input, *credentials)
	enc := json.NewEncoder(os.Stdout)
	if err = enc.Encode(result); err != nil {
		return err
	}
	return nil
}

func SignRequest(request cli.Request, credentials Credentials) cli.Request {
	result := cli.Request(request)
	headers := http.Header(request.Header)
	now := time.Now()
	headers.Set("X-Ovh-Application", credentials.ClientId)
	headers.Set("X-Ovh-Consumer", credentials.ConsumerKey)
	headers.Set("X-Ovh-Timestamp", fmt.Sprint(now.Unix()))
	headers.Set("X-Ovh-Signature", SignPayload(credentials, result.Method, result.URI, result.Body, now))
	return result
}

func SignPayload(credentials Credentials, method string, url string, body string, timestamp time.Time) string {
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

type Credentials struct {
	ClientId     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	ConsumerKey  string `json:"consumerKey,omitempty"`
}
