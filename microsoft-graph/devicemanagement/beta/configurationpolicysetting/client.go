package configurationpolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationPolicySettingClient struct {
	Client *msgraph.Client
}

func NewConfigurationPolicySettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationPolicySettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationpolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationPolicySettingClient: %+v", err)
	}

	return &ConfigurationPolicySettingClient{
		Client: client,
	}, nil
}
