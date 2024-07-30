package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtection struct {
	Bitlocker                  *Bitlocker                   `json:"bitlocker,omitempty"`
	DataLossPreventionPolicies *[]DataLossPreventionPolicy  `json:"dataLossPreventionPolicies,omitempty"`
	Id                         *string                      `json:"id,omitempty"`
	ODataType                  *string                      `json:"@odata.type,omitempty"`
	Policy                     *InformationProtectionPolicy `json:"policy,omitempty"`
	SensitivityLabels          *[]SensitivityLabel          `json:"sensitivityLabels,omitempty"`
	SensitivityPolicySettings  *SensitivityPolicySettings   `json:"sensitivityPolicySettings,omitempty"`
	ThreatAssessmentRequests   *[]ThreatAssessmentRequest   `json:"threatAssessmentRequests,omitempty"`
}
