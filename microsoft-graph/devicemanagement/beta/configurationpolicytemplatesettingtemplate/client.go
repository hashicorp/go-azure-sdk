package configurationpolicytemplatesettingtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationPolicyTemplateSettingTemplateClient struct {
	Client *msgraph.Client
}

func NewConfigurationPolicyTemplateSettingTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationPolicyTemplateSettingTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationpolicytemplatesettingtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationPolicyTemplateSettingTemplateClient: %+v", err)
	}

	return &ConfigurationPolicyTemplateSettingTemplateClient{
		Client: client,
	}, nil
}
