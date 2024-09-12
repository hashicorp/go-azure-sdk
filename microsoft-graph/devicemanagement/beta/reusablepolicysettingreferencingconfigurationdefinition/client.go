package reusablepolicysettingreferencingconfigurationdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReusablePolicySettingReferencingConfigurationDefinitionClient struct {
	Client *msgraph.Client
}

func NewReusablePolicySettingReferencingConfigurationDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ReusablePolicySettingReferencingConfigurationDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "reusablepolicysettingreferencingconfigurationdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReusablePolicySettingReferencingConfigurationDefinitionClient: %+v", err)
	}

	return &ReusablePolicySettingReferencingConfigurationDefinitionClient{
		Client: client,
	}, nil
}
