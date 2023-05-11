package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"
	"gitlab.com/girlfag/vault-plugin-secrets-surrealdb/pkg"
)

func main() {
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := apiClientMeta.GetTLSConfig()

	err := plugin.ServeMultiplex(&plugin.ServeOpts{
		BackendFactoryFunc: pkg.Factory,
		TLSProviderFunc:    api.VaultPluginTLSProvider(tlsConfig),
	})
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}
