package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyLabelAction struct {
	ActionSource                *ApplyLabelActionActionSource  `json:"actionSource,omitempty"`
	Actions                     *[]InformationProtectionAction `json:"actions,omitempty"`
	Label                       *LabelDetails                  `json:"label,omitempty"`
	ODataType                   *string                        `json:"@odata.type,omitempty"`
	ResponsibleSensitiveTypeIds *[]string                      `json:"responsibleSensitiveTypeIds,omitempty"`
}
