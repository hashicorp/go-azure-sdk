package userexperienceanalyticsbatteryhealthmodelperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthModelPerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBatteryHealthModelPerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBatteryHealthModelPerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbatteryhealthmodelperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBatteryHealthModelPerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsBatteryHealthModelPerformanceClient{
		Client: client,
	}, nil
}
