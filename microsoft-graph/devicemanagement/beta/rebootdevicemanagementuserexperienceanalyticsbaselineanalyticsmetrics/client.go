package rebootdevicemanagementuserexperienceanalyticsbaselineanalyticsmetrics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClient struct {
	Client *msgraph.Client
}

func NewRebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClientWithBaseURI(sdkApi sdkEnv.Api) (*RebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rebootdevicemanagementuserexperienceanalyticsbaselineanalyticsmetrics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClient: %+v", err)
	}

	return &RebootDeviceManagementUserExperienceAnalyticsBaselineAnalyticsMetricsClient{
		Client: client,
	}, nil
}
