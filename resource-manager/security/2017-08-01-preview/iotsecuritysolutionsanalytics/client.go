package iotsecuritysolutionsanalytics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecuritySolutionsAnalyticsClient struct {
	Client *resourcemanager.Client
}

func NewIoTSecuritySolutionsAnalyticsClientWithBaseURI(sdkApi sdkEnv.Api) (*IoTSecuritySolutionsAnalyticsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "iotsecuritysolutionsanalytics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IoTSecuritySolutionsAnalyticsClient: %+v", err)
	}

	return &IoTSecuritySolutionsAnalyticsClient{
		Client: client,
	}, nil
}
