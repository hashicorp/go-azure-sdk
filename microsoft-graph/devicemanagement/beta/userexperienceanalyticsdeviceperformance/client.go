package userexperienceanalyticsdeviceperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDevicePerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDevicePerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDevicePerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdeviceperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDevicePerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsDevicePerformanceClient{
		Client: client,
	}, nil
}
