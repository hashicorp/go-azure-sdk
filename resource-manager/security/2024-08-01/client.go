package v2024_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2024-08-01/customrecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2024-08-01/securitystandards"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2024-08-01/standardassignments"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CustomRecommendations *customrecommendations.CustomRecommendationsClient
	SecurityStandards     *securitystandards.SecurityStandardsClient
	StandardAssignments   *standardassignments.StandardAssignmentsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	customRecommendationsClient, err := customrecommendations.NewCustomRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomRecommendations client: %+v", err)
	}
	configureFunc(customRecommendationsClient.Client)

	securityStandardsClient, err := securitystandards.NewSecurityStandardsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityStandards client: %+v", err)
	}
	configureFunc(securityStandardsClient.Client)

	standardAssignmentsClient, err := standardassignments.NewStandardAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StandardAssignments client: %+v", err)
	}
	configureFunc(standardAssignmentsClient.Client)

	return &Client{
		CustomRecommendations: customRecommendationsClient,
		SecurityStandards:     securityStandardsClient,
		StandardAssignments:   standardAssignmentsClient,
	}, nil
}
