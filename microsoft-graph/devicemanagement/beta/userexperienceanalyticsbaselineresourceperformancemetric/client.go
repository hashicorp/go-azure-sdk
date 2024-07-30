package userexperienceanalyticsbaselineresourceperformancemetric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaselineResourcePerformanceMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBaselineResourcePerformanceMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBaselineResourcePerformanceMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbaselineresourceperformancemetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBaselineResourcePerformanceMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsBaselineResourcePerformanceMetricClient{
		Client: client,
	}, nil
}
