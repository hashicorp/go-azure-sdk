package userexperienceanalyticsworkfromanywherehardwarereadinessmetric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsworkfromanywherehardwarereadinessmetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricClient{
		Client: client,
	}, nil
}
