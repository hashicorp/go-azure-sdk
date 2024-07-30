package userexperienceanalyticsdevicemetrichistory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceMetricHistoryClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDeviceMetricHistoryClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDeviceMetricHistoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdevicemetrichistory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDeviceMetricHistoryClient: %+v", err)
	}

	return &UserExperienceAnalyticsDeviceMetricHistoryClient{
		Client: client,
	}, nil
}
