package userexperienceanalyticsdevicescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDeviceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDeviceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdevicescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDeviceScopeClient: %+v", err)
	}

	return &UserExperienceAnalyticsDeviceScopeClient{
		Client: client,
	}, nil
}
