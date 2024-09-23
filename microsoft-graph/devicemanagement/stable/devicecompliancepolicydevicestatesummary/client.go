package devicecompliancepolicydevicestatesummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyDeviceStateSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyDeviceStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyDeviceStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicydevicestatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyDeviceStateSummaryClient: %+v", err)
	}

	return &DeviceCompliancePolicyDeviceStateSummaryClient{
		Client: client,
	}, nil
}
