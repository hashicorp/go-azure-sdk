package userexperienceanalyticsscorehistory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsScoreHistoryClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsScoreHistoryClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsScoreHistoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsscorehistory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsScoreHistoryClient: %+v", err)
	}

	return &UserExperienceAnalyticsScoreHistoryClient{
		Client: client,
	}, nil
}
