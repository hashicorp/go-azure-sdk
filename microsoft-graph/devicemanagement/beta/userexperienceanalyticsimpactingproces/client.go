package userexperienceanalyticsimpactingproces

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsImpactingProcesClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsImpactingProcesClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsImpactingProcesClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsimpactingproces", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsImpactingProcesClient: %+v", err)
	}

	return &UserExperienceAnalyticsImpactingProcesClient{
		Client: client,
	}, nil
}
