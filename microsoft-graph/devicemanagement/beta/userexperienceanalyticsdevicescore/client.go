package userexperienceanalyticsdevicescore

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScoreClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDeviceScoreClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDeviceScoreClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdevicescore", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDeviceScoreClient: %+v", err)
	}

	return &UserExperienceAnalyticsDeviceScoreClient{
		Client: client,
	}, nil
}
