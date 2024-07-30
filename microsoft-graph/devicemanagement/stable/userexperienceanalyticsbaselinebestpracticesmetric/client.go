package userexperienceanalyticsbaselinebestpracticesmetric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaselineBestPracticesMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBaselineBestPracticesMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBaselineBestPracticesMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbaselinebestpracticesmetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBaselineBestPracticesMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsBaselineBestPracticesMetricClient{
		Client: client,
	}, nil
}
