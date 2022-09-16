package v2022_05_04

import (
	"github.com/Azure/go-autorest/autorest"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	bestPracticesClient := bestpractices.NewBestPracticesClientWithBaseURI(endpoint)
	configureAuthFunc(&bestPracticesClient.Client)

	bestPracticesVersionsClient := bestpracticesversions.NewBestPracticesVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&bestPracticesVersionsClient.Client)

	configurationProfileAssignmentsClient := configurationprofileassignments.NewConfigurationProfileAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationProfileAssignmentsClient.Client)

	configurationProfileHCIAssignmentsClient := configurationprofilehciassignments.NewConfigurationProfileHCIAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationProfileHCIAssignmentsClient.Client)

	configurationProfileHCRPAssignmentsClient := configurationprofilehcrpassignments.NewConfigurationProfileHCRPAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationProfileHCRPAssignmentsClient.Client)

	configurationProfilesClient := configurationprofiles.NewConfigurationProfilesClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationProfilesClient.Client)

	configurationProfilesVersionsClient := configurationprofilesversions.NewConfigurationProfilesVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationProfilesVersionsClient.Client)

	hCIReportsClient := hcireports.NewHCIReportsClientWithBaseURI(endpoint)
	configureAuthFunc(&hCIReportsClient.Client)

	hCRPReportsClient := hcrpreports.NewHCRPReportsClientWithBaseURI(endpoint)
	configureAuthFunc(&hCRPReportsClient.Client)

	reportsClient := reports.NewReportsClientWithBaseURI(endpoint)
	configureAuthFunc(&reportsClient.Client)

	servicePrincipalsClient := serviceprincipals.NewServicePrincipalsClientWithBaseURI(endpoint)
	configureAuthFunc(&servicePrincipalsClient.Client)

	return Client{
		BestPractices:                       &bestPracticesClient,
		BestPracticesVersions:               &bestPracticesVersionsClient,
		ConfigurationProfileAssignments:     &configurationProfileAssignmentsClient,
		ConfigurationProfileHCIAssignments:  &configurationProfileHCIAssignmentsClient,
		ConfigurationProfileHCRPAssignments: &configurationProfileHCRPAssignmentsClient,
		ConfigurationProfiles:               &configurationProfilesClient,
		ConfigurationProfilesVersions:       &configurationProfilesVersionsClient,
		HCIReports:                          &hCIReportsClient,
		HCRPReports:                         &hCRPReportsClient,
		Reports:                             &reportsClient,
		ServicePrincipals:                   &servicePrincipalsClient,
	}
}
