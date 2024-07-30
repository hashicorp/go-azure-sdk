package userexperienceanalyticsnotautopilotreadydevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsNotAutopilotReadyDeviceClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsNotAutopilotReadyDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsNotAutopilotReadyDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsnotautopilotreadydevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsNotAutopilotReadyDeviceClient: %+v", err)
	}

	return &UserExperienceAnalyticsNotAutopilotReadyDeviceClient{
		Client: client,
	}, nil
}
