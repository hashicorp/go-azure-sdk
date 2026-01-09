package userexperienceanalyticscategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsCategoryClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticscategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsCategoryClient: %+v", err)
	}

	return &UserExperienceAnalyticsCategoryClient{
		Client: client,
	}, nil
}
