package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityApplyLabelAction struct {
	ActionSource                *SecurityApplyLabelActionActionSource  `json:"actionSource,omitempty"`
	Actions                     *[]SecurityInformationProtectionAction `json:"actions,omitempty"`
	ODataType                   *string                                `json:"@odata.type,omitempty"`
	ResponsibleSensitiveTypeIds *[]string                              `json:"responsibleSensitiveTypeIds,omitempty"`
	SensitivityLabelId          *string                                `json:"sensitivityLabelId,omitempty"`
}
