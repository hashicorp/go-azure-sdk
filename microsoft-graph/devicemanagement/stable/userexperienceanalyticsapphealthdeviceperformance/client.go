package userexperienceanalyticsapphealthdeviceperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDevicePerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthDevicePerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthDevicePerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthdeviceperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthDevicePerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthDevicePerformanceClient{
		Client: client,
	}, nil
}
