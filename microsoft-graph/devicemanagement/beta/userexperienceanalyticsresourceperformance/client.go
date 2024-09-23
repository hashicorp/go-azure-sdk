package userexperienceanalyticsresourceperformance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsResourcePerformanceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsResourcePerformanceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsResourcePerformanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsresourceperformance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsResourcePerformanceClient: %+v", err)
	}

	return &UserExperienceAnalyticsResourcePerformanceClient{
		Client: client,
	}, nil
}
