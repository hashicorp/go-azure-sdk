package userexperienceanalyticsbatteryhealthosperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthOsPerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBatteryHealthOsPerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBatteryHealthOsPerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbatteryhealthosperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBatteryHealthOsPerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsBatteryHealthOsPerformanceClient{
		Client: client,
	}, nil
}
