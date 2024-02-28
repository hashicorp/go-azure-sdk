package v2019_01_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/alertssuppressionrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/assessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/assessmentsmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/automations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/regulatorycompliance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01-preview/subassessments"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AlertsSuppressionRules *alertssuppressionrules.AlertsSuppressionRulesClient
	Assessments            *assessments.AssessmentsClient
	AssessmentsMetadata    *assessmentsmetadata.AssessmentsMetadataClient
	Automations            *automations.AutomationsClient
	RegulatoryCompliance   *regulatorycompliance.RegulatoryComplianceClient
	SubAssessments         *subassessments.SubAssessmentsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertsSuppressionRulesClient, err := alertssuppressionrules.NewAlertsSuppressionRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertsSuppressionRules client: %+v", err)
	}
	configureFunc(alertsSuppressionRulesClient.Client)

	assessmentsClient, err := assessments.NewAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Assessments client: %+v", err)
	}
	configureFunc(assessmentsClient.Client)

	assessmentsMetadataClient, err := assessmentsmetadata.NewAssessmentsMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AssessmentsMetadata client: %+v", err)
	}
	configureFunc(assessmentsMetadataClient.Client)

	automationsClient, err := automations.NewAutomationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Automations client: %+v", err)
	}
	configureFunc(automationsClient.Client)

	regulatoryComplianceClient, err := regulatorycompliance.NewRegulatoryComplianceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegulatoryCompliance client: %+v", err)
	}
	configureFunc(regulatoryComplianceClient.Client)

	subAssessmentsClient, err := subassessments.NewSubAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubAssessments client: %+v", err)
	}
	configureFunc(subAssessmentsClient.Client)

	return &Client{
		AlertsSuppressionRules: alertsSuppressionRulesClient,
		Assessments:            assessmentsClient,
		AssessmentsMetadata:    assessmentsMetadataClient,
		Automations:            automationsClient,
		RegulatoryCompliance:   regulatoryComplianceClient,
		SubAssessments:         subAssessmentsClient,
	}, nil
}
