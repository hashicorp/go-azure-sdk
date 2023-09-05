package v2022_05_04

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/bestpractices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/bestpracticesversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/configurationprofileassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/configurationprofilehciassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/configurationprofilehcrpassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/configurationprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/configurationprofilesversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/hcireports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/hcrpreports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/reports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/serviceprincipals"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BestPractices                       *bestpractices.BestPracticesClient
	BestPracticesVersions               *bestpracticesversions.BestPracticesVersionsClient
	ConfigurationProfileAssignments     *configurationprofileassignments.ConfigurationProfileAssignmentsClient
	ConfigurationProfileHCIAssignments  *configurationprofilehciassignments.ConfigurationProfileHCIAssignmentsClient
	ConfigurationProfileHCRPAssignments *configurationprofilehcrpassignments.ConfigurationProfileHCRPAssignmentsClient
	ConfigurationProfiles               *configurationprofiles.ConfigurationProfilesClient
	ConfigurationProfilesVersions       *configurationprofilesversions.ConfigurationProfilesVersionsClient
	HCIReports                          *hcireports.HCIReportsClient
	HCRPReports                         *hcrpreports.HCRPReportsClient
	Reports                             *reports.ReportsClient
	ServicePrincipals                   *serviceprincipals.ServicePrincipalsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	bestPracticesClient, err := bestpractices.NewBestPracticesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BestPractices client: %+v", err)
	}
	configureFunc(bestPracticesClient.Client)

	bestPracticesVersionsClient, err := bestpracticesversions.NewBestPracticesVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BestPracticesVersions client: %+v", err)
	}
	configureFunc(bestPracticesVersionsClient.Client)

	configurationProfileAssignmentsClient, err := configurationprofileassignments.NewConfigurationProfileAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationProfileAssignments client: %+v", err)
	}
	configureFunc(configurationProfileAssignmentsClient.Client)

	configurationProfileHCIAssignmentsClient, err := configurationprofilehciassignments.NewConfigurationProfileHCIAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationProfileHCIAssignments client: %+v", err)
	}
	configureFunc(configurationProfileHCIAssignmentsClient.Client)

	configurationProfileHCRPAssignmentsClient, err := configurationprofilehcrpassignments.NewConfigurationProfileHCRPAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationProfileHCRPAssignments client: %+v", err)
	}
	configureFunc(configurationProfileHCRPAssignmentsClient.Client)

	configurationProfilesClient, err := configurationprofiles.NewConfigurationProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationProfiles client: %+v", err)
	}
	configureFunc(configurationProfilesClient.Client)

	configurationProfilesVersionsClient, err := configurationprofilesversions.NewConfigurationProfilesVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationProfilesVersions client: %+v", err)
	}
	configureFunc(configurationProfilesVersionsClient.Client)

	hCIReportsClient, err := hcireports.NewHCIReportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HCIReports client: %+v", err)
	}
	configureFunc(hCIReportsClient.Client)

	hCRPReportsClient, err := hcrpreports.NewHCRPReportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HCRPReports client: %+v", err)
	}
	configureFunc(hCRPReportsClient.Client)

	reportsClient, err := reports.NewReportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Reports client: %+v", err)
	}
	configureFunc(reportsClient.Client)

	servicePrincipalsClient, err := serviceprincipals.NewServicePrincipalsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServicePrincipals client: %+v", err)
	}
	configureFunc(servicePrincipalsClient.Client)

	return &Client{
		BestPractices:                       bestPracticesClient,
		BestPracticesVersions:               bestPracticesVersionsClient,
		ConfigurationProfileAssignments:     configurationProfileAssignmentsClient,
		ConfigurationProfileHCIAssignments:  configurationProfileHCIAssignmentsClient,
		ConfigurationProfileHCRPAssignments: configurationProfileHCRPAssignmentsClient,
		ConfigurationProfiles:               configurationProfilesClient,
		ConfigurationProfilesVersions:       configurationProfilesVersionsClient,
		HCIReports:                          hCIReportsClient,
		HCRPReports:                         hCRPReportsClient,
		Reports:                             reportsClient,
		ServicePrincipals:                   servicePrincipalsClient,
	}, nil
}
