package configurationpolicysettingsettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationPolicySettingSettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewConfigurationPolicySettingSettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationPolicySettingSettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationpolicysettingsettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationPolicySettingSettingDefinitionClient: %+v", err)
	}

	return &ConfigurationPolicySettingSettingDefinitionClient{
		Client: client,
	}, nil
}
