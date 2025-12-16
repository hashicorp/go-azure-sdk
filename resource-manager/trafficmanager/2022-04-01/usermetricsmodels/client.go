package usermetricsmodels

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMetricsModelsClient struct {
	Client *resourcemanager.Client
}

func NewUserMetricsModelsClientWithBaseURI(sdkApi sdkEnv.Api) (*UserMetricsModelsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "usermetricsmodels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMetricsModelsClient: %+v", err)
	}

	return &UserMetricsModelsClient{
		Client: client,
	}, nil
}
