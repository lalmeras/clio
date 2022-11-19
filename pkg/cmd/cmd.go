package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/ovh/go-ovh/ovh"
	"github.com/spf13/cobra"
	"golang.org/x/sync/semaphore"
	"gopkg.in/ini.v1"

	"github.com/lalmeras/clio/pkg/types"
	"github.com/lalmeras/clio/pkg/util"
)

const LOGIN_PORT = 8000

var rootCmd = &cobra.Command{
	Use:   "clio",
	Short: "Clio is a cli OVH API client",
	Long:  `Basic API client for OVH API (https://api.ovh.com/console/#/)`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(meCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(cloudCmd)

	rootCmd.AddGroup(authGroup)
	rootCmd.AddGroup(cloudGroup)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Perform API login",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewDefaultClient()
		if err != nil {
			fmt.Printf("Error: %q\n", err)
			return
		}
		req := client.NewCkRequest()
		req.AddRecursiveRules(ovh.ReadOnly, "/")
		req.Redirection = fmt.Sprintf("http://localhost:%d/", LOGIN_PORT)
		pendingCk, err := req.Do()
		if err != nil {
			fmt.Printf("Error: %q\n", err)
			return
		}
		sem := semaphore.NewWeighted(1)
		sem.Acquire(context.TODO(), 1)
		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", LOGIN_PORT),
			Handler:        HttpHandler{sem},
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		fmt.Printf("Visit %s to validate authentication\n", pendingCk.ValidationURL)
		go s.ListenAndServe()
		sem.Acquire(context.TODO(), 1)
		fmt.Printf("Login valided\n")
		consumerKey := pendingCk.ConsumerKey
		home, err := homedir.Dir()
		ovhconf := path.Join(home, ".ovh.conf")
		cfg, err := ini.Load(ovhconf)
		cfg.Section("ovh-eu").Key("consumer_key").SetValue(consumerKey)
		cfg.SaveTo(ovhconf)
	},
}

type HttpHandler struct {
	sem *semaphore.Weighted
}

func (self HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		self.sem.Release(1)
	}
}

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get details about your nichandle.",
	Run: func(cmd *cobra.Command, args []string) {
		nichandle := types.Nichandle{}
		util.Check(util.OvhGet("/me", &nichandle))
		fmt.Printf("%s\n", nichandle.Email)
	},
}
