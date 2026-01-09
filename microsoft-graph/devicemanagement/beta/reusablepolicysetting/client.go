package reusablepolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReusablePolicySettingClient struct {
	Client *msgraph.Client
}

func NewReusablePolicySettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ReusablePolicySettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "reusablepolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReusablePolicySettingClient: %+v", err)
	}

	return &ReusablePolicySettingClient{
		Client: client,
	}, nil
}
