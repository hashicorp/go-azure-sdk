package devicecompliancepolicysettingstatesummarysettingstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicySettingStateSummarySettingStateClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicySettingStateSummarySettingStateClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicySettingStateSummarySettingStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicysettingstatesummarysettingstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicySettingStateSummarySettingStateClient: %+v", err)
	}

	return &DeviceCompliancePolicySettingStateSummarySettingStateClient{
		Client: client,
	}, nil
}
