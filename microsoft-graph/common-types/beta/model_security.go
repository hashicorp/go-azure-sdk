package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Security struct {
	Alerts                     *[]Alert                       `json:"alerts,omitempty"`
	Alertsv2                   *[]SecurityAlert               `json:"alerts_v2,omitempty"`
	AttackSimulation           *AttackSimulationRoot          `json:"attackSimulation,omitempty"`
	Cases                      *SecurityCasesRoot             `json:"cases,omitempty"`
	CloudAppSecurityProfiles   *[]CloudAppSecurityProfile     `json:"cloudAppSecurityProfiles,omitempty"`
	DomainSecurityProfiles     *[]DomainSecurityProfile       `json:"domainSecurityProfiles,omitempty"`
	FileSecurityProfiles       *[]FileSecurityProfile         `json:"fileSecurityProfiles,omitempty"`
	HostSecurityProfiles       *[]HostSecurityProfile         `json:"hostSecurityProfiles,omitempty"`
	Id                         *string                        `json:"id,omitempty"`
	Incidents                  *[]SecurityIncident            `json:"incidents,omitempty"`
	InformationProtection      *SecurityInformationProtection `json:"informationProtection,omitempty"`
	IpSecurityProfiles         *[]IpSecurityProfile           `json:"ipSecurityProfiles,omitempty"`
	Labels                     *SecurityLabelsRoot            `json:"labels,omitempty"`
	ODataType                  *string                        `json:"@odata.type,omitempty"`
	ProviderStatus             *[]SecurityProviderStatus      `json:"providerStatus,omitempty"`
	ProviderTenantSettings     *[]ProviderTenantSetting       `json:"providerTenantSettings,omitempty"`
	SecureScoreControlProfiles *[]SecureScoreControlProfile   `json:"secureScoreControlProfiles,omitempty"`
	SecureScores               *[]SecureScore                 `json:"secureScores,omitempty"`
	SecurityActions            *[]SecurityAction              `json:"securityActions,omitempty"`
	SubjectRightsRequests      *[]SubjectRightsRequest        `json:"subjectRightsRequests,omitempty"`
	ThreatIntelligence         *SecurityThreatIntelligence    `json:"threatIntelligence,omitempty"`
	ThreatSubmission           *SecurityThreatSubmissionRoot  `json:"threatSubmission,omitempty"`
	TiIndicators               *[]TiIndicator                 `json:"tiIndicators,omitempty"`
	TriggerTypes               *SecurityTriggerTypesRoot      `json:"triggerTypes,omitempty"`
	Triggers                   *SecurityTriggersRoot          `json:"triggers,omitempty"`
	UserSecurityProfiles       *[]UserSecurityProfile         `json:"userSecurityProfiles,omitempty"`
}
