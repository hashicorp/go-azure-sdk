package userexperienceanalyticsanomaly

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAnomalyClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAnomalyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsanomaly", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAnomalyClient: %+v", err)
	}

	return &UserExperienceAnalyticsAnomalyClient{
		Client: client,
	}, nil
}
