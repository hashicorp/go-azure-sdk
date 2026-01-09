package userexperienceanalyticsapphealthapplicationperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthApplicationPerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthApplicationPerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthApplicationPerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthapplicationperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthApplicationPerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthApplicationPerformanceClient{
		Client: client,
	}, nil
}
