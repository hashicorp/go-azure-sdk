package v2020_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/adaptivenetworkhardenings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/allowedconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/applicationwhitelistings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/assessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/assessmentsmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/discoveredsecuritysolutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/externalsecuritysolutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/jitnetworkaccesspolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescorecontroldefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securescorecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securitysolutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securitysolutionsreferencedata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/servervulnerabilityassessment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/servervulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/topology"
)

type Client struct {
	AdaptiveNetworkHardenings      *adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient
	Alerts                         *alerts.AlertsClient
	AllowedConnections             *allowedconnections.AllowedConnectionsClient
	ApplicationWhitelistings       *applicationwhitelistings.ApplicationWhitelistingsClient
	Assessments                    *assessments.AssessmentsClient
	AssessmentsMetadata            *assessmentsmetadata.AssessmentsMetadataClient
	DiscoveredSecuritySolutions    *discoveredsecuritysolutions.DiscoveredSecuritySolutionsClient
	ExternalSecuritySolutions      *externalsecuritysolutions.ExternalSecuritySolutionsClient
	JitNetworkAccessPolicies       *jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient
	SecureScore                    *securescore.SecureScoreClient
	SecureScoreControlDefinitions  *securescorecontroldefinitions.SecureScoreControlDefinitionsClient
	SecureScoreControls            *securescorecontrols.SecureScoreControlsClient
	SecuritySolutions              *securitysolutions.SecuritySolutionsClient
	SecuritySolutionsReferenceData *securitysolutionsreferencedata.SecuritySolutionsReferenceDataClient
	ServerVulnerabilityAssessment  *servervulnerabilityassessment.ServerVulnerabilityAssessmentClient
	ServerVulnerabilityAssessments *servervulnerabilityassessments.ServerVulnerabilityAssessmentsClient
	Topology                       *topology.TopologyClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	adaptiveNetworkHardeningsClient := adaptivenetworkhardenings.NewAdaptiveNetworkHardeningsClientWithBaseURI(endpoint)
	configureAuthFunc(&adaptiveNetworkHardeningsClient.Client)

	alertsClient := alerts.NewAlertsClientWithBaseURI(endpoint)
	configureAuthFunc(&alertsClient.Client)

	allowedConnectionsClient := allowedconnections.NewAllowedConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&allowedConnectionsClient.Client)

	applicationWhitelistingsClient := applicationwhitelistings.NewApplicationWhitelistingsClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationWhitelistingsClient.Client)

	assessmentsClient := assessments.NewAssessmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&assessmentsClient.Client)

	assessmentsMetadataClient := assessmentsmetadata.NewAssessmentsMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&assessmentsMetadataClient.Client)

	discoveredSecuritySolutionsClient := discoveredsecuritysolutions.NewDiscoveredSecuritySolutionsClientWithBaseURI(endpoint)
	configureAuthFunc(&discoveredSecuritySolutionsClient.Client)

	externalSecuritySolutionsClient := externalsecuritysolutions.NewExternalSecuritySolutionsClientWithBaseURI(endpoint)
	configureAuthFunc(&externalSecuritySolutionsClient.Client)

	jitNetworkAccessPoliciesClient := jitnetworkaccesspolicies.NewJitNetworkAccessPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&jitNetworkAccessPoliciesClient.Client)

	secureScoreClient := securescore.NewSecureScoreClientWithBaseURI(endpoint)
	configureAuthFunc(&secureScoreClient.Client)

	secureScoreControlDefinitionsClient := securescorecontroldefinitions.NewSecureScoreControlDefinitionsClientWithBaseURI(endpoint)
	configureAuthFunc(&secureScoreControlDefinitionsClient.Client)

	secureScoreControlsClient := securescorecontrols.NewSecureScoreControlsClientWithBaseURI(endpoint)
	configureAuthFunc(&secureScoreControlsClient.Client)

	securitySolutionsClient := securitysolutions.NewSecuritySolutionsClientWithBaseURI(endpoint)
	configureAuthFunc(&securitySolutionsClient.Client)

	securitySolutionsReferenceDataClient := securitysolutionsreferencedata.NewSecuritySolutionsReferenceDataClientWithBaseURI(endpoint)
	configureAuthFunc(&securitySolutionsReferenceDataClient.Client)

	serverVulnerabilityAssessmentClient := servervulnerabilityassessment.NewServerVulnerabilityAssessmentClientWithBaseURI(endpoint)
	configureAuthFunc(&serverVulnerabilityAssessmentClient.Client)

	serverVulnerabilityAssessmentsClient := servervulnerabilityassessments.NewServerVulnerabilityAssessmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&serverVulnerabilityAssessmentsClient.Client)

	topologyClient := topology.NewTopologyClientWithBaseURI(endpoint)
	configureAuthFunc(&topologyClient.Client)

	return Client{
		AdaptiveNetworkHardenings:      &adaptiveNetworkHardeningsClient,
		Alerts:                         &alertsClient,
		AllowedConnections:             &allowedConnectionsClient,
		ApplicationWhitelistings:       &applicationWhitelistingsClient,
		Assessments:                    &assessmentsClient,
		AssessmentsMetadata:            &assessmentsMetadataClient,
		DiscoveredSecuritySolutions:    &discoveredSecuritySolutionsClient,
		ExternalSecuritySolutions:      &externalSecuritySolutionsClient,
		JitNetworkAccessPolicies:       &jitNetworkAccessPoliciesClient,
		SecureScore:                    &secureScoreClient,
		SecureScoreControlDefinitions:  &secureScoreControlDefinitionsClient,
		SecureScoreControls:            &secureScoreControlsClient,
		SecuritySolutions:              &securitySolutionsClient,
		SecuritySolutionsReferenceData: &securitySolutionsReferenceDataClient,
		ServerVulnerabilityAssessment:  &serverVulnerabilityAssessmentClient,
		ServerVulnerabilityAssessments: &serverVulnerabilityAssessmentsClient,
		Topology:                       &topologyClient,
	}
}
