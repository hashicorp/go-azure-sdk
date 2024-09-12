package userexperienceanalyticsapphealthoverview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthOverviewClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAppHealthOverviewClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAppHealthOverviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsapphealthoverview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAppHealthOverviewClient: %+v", err)
	}

	return &UserExperienceAnalyticsAppHealthOverviewClient{
		Client: client,
	}, nil
}
