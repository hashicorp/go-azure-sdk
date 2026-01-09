package userexperienceanalyticsbatteryhealthruntimedetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthRuntimeDetailClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBatteryHealthRuntimeDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBatteryHealthRuntimeDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbatteryhealthruntimedetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBatteryHealthRuntimeDetailClient: %+v", err)
	}

	return &UserExperienceAnalyticsBatteryHealthRuntimeDetailClient{
		Client: client,
	}, nil
}
