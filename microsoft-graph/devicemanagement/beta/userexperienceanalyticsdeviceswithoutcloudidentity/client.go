package userexperienceanalyticsdeviceswithoutcloudidentity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDevicesWithoutCloudIdentityClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDevicesWithoutCloudIdentityClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDevicesWithoutCloudIdentityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdeviceswithoutcloudidentity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDevicesWithoutCloudIdentityClient: %+v", err)
	}

	return &UserExperienceAnalyticsDevicesWithoutCloudIdentityClient{
		Client: client,
	}, nil
}
