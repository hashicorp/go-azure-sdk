package configurationpolicytemplatesettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationPolicyTemplateSettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewConfigurationPolicyTemplateSettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationPolicyTemplateSettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationpolicytemplatesettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationPolicyTemplateSettingDefinitionClient: %+v", err)
	}

	return &ConfigurationPolicyTemplateSettingDefinitionClient{
		Client: client,
	}, nil
}
