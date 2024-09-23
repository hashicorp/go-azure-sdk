package userexperienceanalyticsbatteryhealthdeviceperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDevicePerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBatteryHealthDevicePerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBatteryHealthDevicePerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbatteryhealthdeviceperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBatteryHealthDevicePerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsBatteryHealthDevicePerformanceClient{
		Client: client,
	}, nil
}
