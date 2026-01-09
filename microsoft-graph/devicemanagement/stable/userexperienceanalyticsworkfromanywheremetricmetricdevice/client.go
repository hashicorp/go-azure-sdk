package userexperienceanalyticsworkfromanywheremetricmetricdevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsworkfromanywheremetricmetricdevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient: %+v", err)
	}

	return &UserExperienceAnalyticsWorkFromAnywhereMetricMetricDeviceClient{
		Client: client,
	}, nil
}
