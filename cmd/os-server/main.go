package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	sdk "github.com/configcat/go-sdk/v9"
	configcat "github.com/open-feature/go-sdk-contrib/providers/configcat/pkg"
	"github.com/open-feature/go-sdk/openfeature"
)

func main() {
	SDKKey, ok := os.LookupEnv("CONFIGCAT_SDK_KEY")
	if !ok {
		log.Fatal("CONFIGCAT_SDK_KEY is not set")
	}

	sdkClient := sdk.NewCustomClient(sdk.Config{
		SDKKey:       SDKKey,
		PollInterval: 60 * time.Second,
		FlagOverrides: &sdk.FlagOverrides{
			FilePath: "./local_flags.json",
			Behavior: sdk.LocalOverRemote,
		},
	})
	defer sdkClient.Close()

	provider := configcat.NewProvider(sdkClient)
	openfeature.SetProvider(provider)
	ccClient := openfeature.NewClient("app")

	srv := NewServer(ccClient)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
