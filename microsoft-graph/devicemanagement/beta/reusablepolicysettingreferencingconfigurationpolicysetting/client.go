package reusablepolicysettingreferencingconfigurationpolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReusablePolicySettingReferencingConfigurationPolicySettingClient struct {
	Client *msgraph.Client
}

func NewReusablePolicySettingReferencingConfigurationPolicySettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ReusablePolicySettingReferencingConfigurationPolicySettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "reusablepolicysettingreferencingconfigurationpolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReusablePolicySettingReferencingConfigurationPolicySettingClient: %+v", err)
	}

	return &ReusablePolicySettingReferencingConfigurationPolicySettingClient{
		Client: client,
	}, nil
}
