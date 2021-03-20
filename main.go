package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/shekhar-jha/terraform-provider-env/provider"
	"log"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	commit string = ""
)

func main() {
	log.Println("Starting 'env' provider...")
	defer log.Println("Exiting 'env' provider.")
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()
	log.Printf("Debug mode %t\n", debugMode)

	opts := &plugin.ServeOpts{ProviderFunc: providerSchema(version)}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/shekhar-jha/env", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
	log.Println("Start serving....")
	plugin.Serve(opts)
}

func providerSchema(version string) func() *schema.Provider {
	log.Printf("Entering providerSchema(version: %s)", version)
	defer log.Printf("Exiting providerSchema(version)")
	return func() *schema.Provider {
		log.Println("Entering providerSchema.func()")
		defer log.Println("Exiting providerSchema.func()")
		return provider.Provider()
	}
}
