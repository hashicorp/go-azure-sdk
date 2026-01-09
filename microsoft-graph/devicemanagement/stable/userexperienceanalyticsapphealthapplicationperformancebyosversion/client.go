package userexperienceanalyticsapphealthapplicationperformancebyosversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthapplicationperformancebyosversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionClient{
		Client: client,
	}, nil
}
