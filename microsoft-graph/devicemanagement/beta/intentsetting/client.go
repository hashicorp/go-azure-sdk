package intentsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentSettingClient struct {
	Client *msgraph.Client
}

func NewIntentSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentSettingClient: %+v", err)
	}

	return &IntentSettingClient{
		Client: client,
	}, nil
}
