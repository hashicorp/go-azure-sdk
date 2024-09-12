package userexperienceanalyticsworkfromanywheremetricdevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereMetricDeviceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsWorkFromAnywhereMetricDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsWorkFromAnywhereMetricDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsworkfromanywheremetricdevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsWorkFromAnywhereMetricDeviceClient: %+v", err)
	}

	return &UserExperienceAnalyticsWorkFromAnywhereMetricDeviceClient{
		Client: client,
	}, nil
}
