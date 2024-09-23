package userexperienceanalyticsapphealthdeviceperformancedetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDevicePerformanceDetailClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthDevicePerformanceDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthDevicePerformanceDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthdeviceperformancedetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthDevicePerformanceDetailClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthDevicePerformanceDetailClient{
		Client: client,
	}, nil
}
