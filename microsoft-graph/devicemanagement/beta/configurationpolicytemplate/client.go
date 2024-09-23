package configurationpolicytemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationPolicyTemplateClient struct {
	Client *msgraph.Client
}

func NewConfigurationPolicyTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationPolicyTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationpolicytemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationPolicyTemplateClient: %+v", err)
	}

	return &ConfigurationPolicyTemplateClient{
		Client: client,
	}, nil
}
