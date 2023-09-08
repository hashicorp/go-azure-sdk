package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Security struct {
	Alerts                     *[]Alert                     `json:"alerts,omitempty"`
	Alertsv2                   *[]SecurityAlert             `json:"alerts_v2,omitempty"`
	AttackSimulation           *AttackSimulationRoot        `json:"attackSimulation,omitempty"`
	Cases                      *SecurityCasesRoot           `json:"cases,omitempty"`
	Id                         *string                      `json:"id,omitempty"`
	Incidents                  *[]SecurityIncident          `json:"incidents,omitempty"`
	ODataType                  *string                      `json:"@odata.type,omitempty"`
	SecureScoreControlProfiles *[]SecureScoreControlProfile `json:"secureScoreControlProfiles,omitempty"`
	SecureScores               *[]SecureScore               `json:"secureScores,omitempty"`
	ThreatIntelligence         *SecurityThreatIntelligence  `json:"threatIntelligence,omitempty"`
	TriggerTypes               *SecurityTriggerTypesRoot    `json:"triggerTypes,omitempty"`
	Triggers                   *SecurityTriggersRoot        `json:"triggers,omitempty"`
}
