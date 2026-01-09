package reusablesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReusableSettingClient struct {
	Client *msgraph.Client
}

func NewReusableSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ReusableSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "reusablesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReusableSettingClient: %+v", err)
	}

	return &ReusableSettingClient{
		Client: client,
	}, nil
}
