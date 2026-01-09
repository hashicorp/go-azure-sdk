package devicecompliancepolicysettingstatesummarydevicecompliancesettingstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicysettingstatesummarydevicecompliancesettingstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient: %+v", err)
	}

	return &DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient{
		Client: client,
	}, nil
}
