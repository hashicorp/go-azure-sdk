package advancedthreatprotectiononboardingstatesummaryadvancedthreatprotectiononboardingdevicesettingstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClient struct {
	Client *msgraph.Client
}

func NewAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClientWithBaseURI(sdkApi sdkEnv.Api) (*AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "advancedthreatprotectiononboardingstatesummaryadvancedthreatprotectiononboardingdevicesettingstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClient: %+v", err)
	}

	return &AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStateClient{
		Client: client,
	}, nil
}
