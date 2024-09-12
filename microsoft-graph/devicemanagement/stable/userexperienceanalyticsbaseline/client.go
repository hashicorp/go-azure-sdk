package userexperienceanalyticsbaseline

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBaselineClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBaselineClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBaselineClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbaseline", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBaselineClient: %+v", err)
	}

	return &UserExperienceAnalyticsBaselineClient{
		Client: client,
	}, nil
}
