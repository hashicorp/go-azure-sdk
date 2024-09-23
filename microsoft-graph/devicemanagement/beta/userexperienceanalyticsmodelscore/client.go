package userexperienceanalyticsmodelscore

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsModelScoreClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsModelScoreClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsModelScoreClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsmodelscore", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsModelScoreClient: %+v", err)
	}

	return &UserExperienceAnalyticsModelScoreClient{
		Client: client,
	}, nil
}
