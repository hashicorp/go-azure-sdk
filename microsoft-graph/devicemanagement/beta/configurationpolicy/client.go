package configurationpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationPolicyClient struct {
	Client *msgraph.Client
}

func NewConfigurationPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationPolicyClient: %+v", err)
	}

	return &ConfigurationPolicyClient{
		Client: client,
	}, nil
}
