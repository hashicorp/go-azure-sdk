package v2019_01_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/alertssuppressionrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/assessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/assessmentsmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/automations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/regulatorycompliance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/subassessments"
)

type Client struct {
	AlertsSuppressionRules *alertssuppressionrules.AlertsSuppressionRulesClient
	Assessments            *assessments.AssessmentsClient
	AssessmentsMetadata    *assessmentsmetadata.AssessmentsMetadataClient
	Automations            *automations.AutomationsClient
	RegulatoryCompliance   *regulatorycompliance.RegulatoryComplianceClient
	SubAssessments         *subassessments.SubAssessmentsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	alertsSuppressionRulesClient := alertssuppressionrules.NewAlertsSuppressionRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&alertsSuppressionRulesClient.Client)

	assessmentsClient := assessments.NewAssessmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&assessmentsClient.Client)

	assessmentsMetadataClient := assessmentsmetadata.NewAssessmentsMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&assessmentsMetadataClient.Client)

	automationsClient := automations.NewAutomationsClientWithBaseURI(endpoint)
	configureAuthFunc(&automationsClient.Client)

	regulatoryComplianceClient := regulatorycompliance.NewRegulatoryComplianceClientWithBaseURI(endpoint)
	configureAuthFunc(&regulatoryComplianceClient.Client)

	subAssessmentsClient := subassessments.NewSubAssessmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&subAssessmentsClient.Client)

	return Client{
		AlertsSuppressionRules: &alertsSuppressionRulesClient,
		Assessments:            &assessmentsClient,
		AssessmentsMetadata:    &assessmentsMetadataClient,
		Automations:            &automationsClient,
		RegulatoryCompliance:   &regulatoryComplianceClient,
		SubAssessments:         &subAssessmentsClient,
	}
}
