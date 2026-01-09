package userexperienceanalyticsdevicetimelineevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceTimelineEventClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDeviceTimelineEventClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDeviceTimelineEventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdevicetimelineevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDeviceTimelineEventClient: %+v", err)
	}

	return &UserExperienceAnalyticsDeviceTimelineEventClient{
		Client: client,
	}, nil
}
