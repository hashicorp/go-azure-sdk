package userexperienceanalyticsoverview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsOverviewClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsOverviewClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsOverviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsoverview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsOverviewClient: %+v", err)
	}

	return &UserExperienceAnalyticsOverviewClient{
		Client: client,
	}, nil
}
