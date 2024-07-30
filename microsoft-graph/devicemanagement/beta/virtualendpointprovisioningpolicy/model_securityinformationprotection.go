package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInformationProtection struct {
	Id                  *string                                     `json:"id,omitempty"`
	LabelPolicySettings *SecurityInformationProtectionPolicySetting `json:"labelPolicySettings,omitempty"`
	ODataType           *string                                     `json:"@odata.type,omitempty"`
	SensitivityLabels   *[]SecuritySensitivityLabel                 `json:"sensitivityLabels,omitempty"`
}
