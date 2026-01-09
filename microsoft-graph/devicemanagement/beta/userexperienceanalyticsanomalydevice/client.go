package userexperienceanalyticsanomalydevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyDeviceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAnomalyDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAnomalyDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsanomalydevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAnomalyDeviceClient: %+v", err)
	}

	return &UserExperienceAnalyticsAnomalyDeviceClient{
		Client: client,
	}, nil
}
