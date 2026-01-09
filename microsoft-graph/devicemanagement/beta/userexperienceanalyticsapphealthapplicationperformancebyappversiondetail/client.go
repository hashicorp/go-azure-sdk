package userexperienceanalyticsapphealthapplicationperformancebyappversiondetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthapplicationperformancebyappversiondetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailClient{
		Client: client,
	}, nil
}
