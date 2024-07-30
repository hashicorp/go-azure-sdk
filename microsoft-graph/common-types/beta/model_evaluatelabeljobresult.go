package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluateLabelJobResult struct {
	ODataType                 *string                     `json:"@odata.type,omitempty"`
	ResponsiblePolicy         *ResponsiblePolicy          `json:"responsiblePolicy,omitempty"`
	ResponsibleSensitiveTypes *[]ResponsibleSensitiveType `json:"responsibleSensitiveTypes,omitempty"`
	SensitivityLabel          *MatchingLabel              `json:"sensitivityLabel,omitempty"`
}
