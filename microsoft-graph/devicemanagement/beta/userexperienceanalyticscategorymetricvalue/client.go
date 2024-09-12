package userexperienceanalyticscategorymetricvalue

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsCategoryMetricValueClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsCategoryMetricValueClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsCategoryMetricValueClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticscategorymetricvalue", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsCategoryMetricValueClient: %+v", err)
	}

	return &UserExperienceAnalyticsCategoryMetricValueClient{
		Client: client,
	}, nil
}
