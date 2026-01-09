package advancedthreatprotectiononboardingstatesummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionOnboardingStateSummaryClient struct {
	Client *msgraph.Client
}

func NewAdvancedThreatProtectionOnboardingStateSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*AdvancedThreatProtectionOnboardingStateSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "advancedthreatprotectiononboardingstatesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdvancedThreatProtectionOnboardingStateSummaryClient: %+v", err)
	}

	return &AdvancedThreatProtectionOnboardingStateSummaryClient{
		Client: client,
	}, nil
}
