package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/configcat/go-sdk/v9"
)

func main() {
	SDKKey, ok := os.LookupEnv("CONFIGCAT_SDK_KEY")
	if !ok {
		log.Fatal("CONFIGCAT_SDK_KEY is not set")
	}

	ccClient := configcat.NewCustomClient(configcat.Config{
		SDKKey:       SDKKey,
		PollInterval: 60 * time.Second,
		FlagOverrides: &configcat.FlagOverrides{
			FilePath: "./local_flags.json",
			Behavior: configcat.LocalOverRemote,
		},
	})
	defer ccClient.Close()

	srv := NewServer(ccClient)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
