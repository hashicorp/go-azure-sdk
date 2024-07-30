package devicecompliancepolicydevicesettingstatesummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyDeviceSettingStateSummaryClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyDeviceSettingStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyDeviceSettingStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicydevicesettingstatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyDeviceSettingStateSummaryClient: %+v", err)
	}

	return &DeviceCompliancePolicyDeviceSettingStateSummaryClient{
		Client: client,
	}, nil
}
