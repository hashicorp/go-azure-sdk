package userexperienceanalyticsapphealthdevicemodelperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthdevicemodelperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthDeviceModelPerformanceClient{
		Client: client,
	}, nil
}
