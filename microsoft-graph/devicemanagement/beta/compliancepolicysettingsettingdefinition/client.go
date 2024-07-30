package compliancepolicysettingsettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompliancePolicySettingSettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewCompliancePolicySettingSettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*CompliancePolicySettingSettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancepolicysettingsettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CompliancePolicySettingSettingDefinitionClient: %+v", err)
	}

	return &CompliancePolicySettingSettingDefinitionClient{
		Client: client,
	}, nil
}
