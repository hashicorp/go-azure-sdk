package configurationpolicysettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationPolicySettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewConfigurationPolicySettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationPolicySettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationpolicysettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationPolicySettingDefinitionClient: %+v", err)
	}

	return &ConfigurationPolicySettingDefinitionClient{
		Client: client,
	}, nil
}
