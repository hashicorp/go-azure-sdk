package configurationpolicytemplatesettingtemplatesettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationPolicyTemplateSettingTemplateSettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewConfigurationPolicyTemplateSettingTemplateSettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationPolicyTemplateSettingTemplateSettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationpolicytemplatesettingtemplatesettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationPolicyTemplateSettingTemplateSettingDefinitionClient: %+v", err)
	}

	return &ConfigurationPolicyTemplateSettingTemplateSettingDefinitionClient{
		Client: client,
	}, nil
}
