package userexperienceanalyticsapphealthosversionperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthOSVersionPerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthOSVersionPerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthOSVersionPerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthosversionperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthOSVersionPerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthOSVersionPerformanceClient{
		Client: client,
	}, nil
}
