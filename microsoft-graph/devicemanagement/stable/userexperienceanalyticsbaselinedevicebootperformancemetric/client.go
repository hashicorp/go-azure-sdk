package userexperienceanalyticsbaselinedevicebootperformancemetric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbaselinedevicebootperformancemetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsBaselineDeviceBootPerformanceMetricClient{
		Client: client,
	}, nil
}
