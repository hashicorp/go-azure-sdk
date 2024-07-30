package reusablepolicysettingreferencingconfigurationpolicysettingsettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "reusablepolicysettingreferencingconfigurationpolicysettingsettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClient: %+v", err)
	}

	return &ReusablePolicySettingReferencingConfigurationPolicySettingSettingDefinitionClient{
		Client: client,
	}, nil
}
