package userexperienceanalyticsbaselinebatteryhealthmetric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaselineBatteryHealthMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBaselineBatteryHealthMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBaselineBatteryHealthMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbaselinebatteryhealthmetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBaselineBatteryHealthMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsBaselineBatteryHealthMetricClient{
		Client: client,
	}, nil
}
