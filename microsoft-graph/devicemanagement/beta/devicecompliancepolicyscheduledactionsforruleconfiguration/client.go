package devicecompliancepolicyscheduledactionsforruleconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyScheduledActionsForRuleConfigurationClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyScheduledActionsForRuleConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyScheduledActionsForRuleConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicyscheduledactionsforruleconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyScheduledActionsForRuleConfigurationClient: %+v", err)
	}

	return &DeviceCompliancePolicyScheduledActionsForRuleConfigurationClient{
		Client: client,
	}, nil
}
