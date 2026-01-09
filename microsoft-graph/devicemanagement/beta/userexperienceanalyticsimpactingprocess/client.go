package userexperienceanalyticsimpactingprocess

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsImpactingProcessClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsImpactingProcessClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsImpactingProcessClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsimpactingprocess", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsImpactingProcessClient: %+v", err)
	}

	return &UserExperienceAnalyticsImpactingProcessClient{
		Client: client,
	}, nil
}
