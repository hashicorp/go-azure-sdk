package userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdClient{
		Client: client,
	}, nil
}
