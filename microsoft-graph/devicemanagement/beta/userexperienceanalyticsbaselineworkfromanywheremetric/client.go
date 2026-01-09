package userexperienceanalyticsbaselineworkfromanywheremetric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaselineWorkFromAnywhereMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBaselineWorkFromAnywhereMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBaselineWorkFromAnywhereMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbaselineworkfromanywheremetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBaselineWorkFromAnywhereMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsBaselineWorkFromAnywhereMetricClient{
		Client: client,
	}, nil
}
