package userexperienceanalyticsbatteryhealthappimpact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthAppImpactClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBatteryHealthAppImpactClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBatteryHealthAppImpactClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbatteryhealthappimpact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBatteryHealthAppImpactClient: %+v", err)
	}

	return &UserExperienceAnalyticsBatteryHealthAppImpactClient{
		Client: client,
	}, nil
}
