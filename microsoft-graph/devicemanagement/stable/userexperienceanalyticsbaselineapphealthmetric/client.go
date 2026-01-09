package userexperienceanalyticsbaselineapphealthmetric

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaselineAppHealthMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBaselineAppHealthMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBaselineAppHealthMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbaselineapphealthmetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBaselineAppHealthMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsBaselineAppHealthMetricClient{
		Client: client,
	}, nil
}
