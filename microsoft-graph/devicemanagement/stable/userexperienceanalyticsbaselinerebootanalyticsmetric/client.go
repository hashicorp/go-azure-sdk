package userexperienceanalyticsbaselinerebootanalyticsmetric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaselineRebootAnalyticsMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBaselineRebootAnalyticsMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBaselineRebootAnalyticsMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbaselinerebootanalyticsmetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBaselineRebootAnalyticsMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsBaselineRebootAnalyticsMetricClient{
		Client: client,
	}, nil
}
