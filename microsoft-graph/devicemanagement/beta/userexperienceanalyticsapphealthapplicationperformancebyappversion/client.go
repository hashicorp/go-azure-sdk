package userexperienceanalyticsapphealthapplicationperformancebyappversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthapplicationperformancebyappversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionClient{
		Client: client,
	}, nil
}
