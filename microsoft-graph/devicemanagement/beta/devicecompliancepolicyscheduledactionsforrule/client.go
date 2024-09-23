package devicecompliancepolicyscheduledactionsforrule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyScheduledActionsForRuleClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyScheduledActionsForRuleClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyScheduledActionsForRuleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicyscheduledactionsforrule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyScheduledActionsForRuleClient: %+v", err)
	}

	return &DeviceCompliancePolicyScheduledActionsForRuleClient{
		Client: client,
	}, nil
}
