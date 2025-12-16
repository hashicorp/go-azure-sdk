package v2022_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2022-04-01/profiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2022-04-01/trafficmanagergeographichierarchies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2022-04-01/trafficmanagers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2022-04-01/usermetricsmodels"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Profiles                            *profiles.ProfilesClient
	TrafficManagerGeographicHierarchies *trafficmanagergeographichierarchies.TrafficManagerGeographicHierarchiesClient
	Trafficmanagers                     *trafficmanagers.TrafficmanagersClient
	UserMetricsModels                   *usermetricsmodels.UserMetricsModelsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	profilesClient, err := profiles.NewProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Profiles client: %+v", err)
	}
	configureFunc(profilesClient.Client)

	trafficManagerGeographicHierarchiesClient, err := trafficmanagergeographichierarchies.NewTrafficManagerGeographicHierarchiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TrafficManagerGeographicHierarchies client: %+v", err)
	}
	configureFunc(trafficManagerGeographicHierarchiesClient.Client)

	trafficmanagersClient, err := trafficmanagers.NewTrafficmanagersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Trafficmanagers client: %+v", err)
	}
	configureFunc(trafficmanagersClient.Client)

	userMetricsModelsClient, err := usermetricsmodels.NewUserMetricsModelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMetricsModels client: %+v", err)
	}
	configureFunc(userMetricsModelsClient.Client)

	return &Client{
		Profiles:                            profilesClient,
		TrafficManagerGeographicHierarchies: trafficManagerGeographicHierarchiesClient,
		Trafficmanagers:                     trafficmanagersClient,
		UserMetricsModels:                   userMetricsModelsClient,
	}, nil
}
