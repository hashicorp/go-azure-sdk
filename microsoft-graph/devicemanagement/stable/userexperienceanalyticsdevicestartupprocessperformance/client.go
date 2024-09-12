package userexperienceanalyticsdevicestartupprocessperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupProcessPerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDeviceStartupProcessPerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDeviceStartupProcessPerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdevicestartupprocessperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDeviceStartupProcessPerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsDeviceStartupProcessPerformanceClient{
		Client: client,
	}, nil
}
