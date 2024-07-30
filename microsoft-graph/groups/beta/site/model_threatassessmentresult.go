package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentResult struct {
	CreatedDateTime *string                           `json:"createdDateTime,omitempty"`
	Id              *string                           `json:"id,omitempty"`
	Message         *string                           `json:"message,omitempty"`
	ODataType       *string                           `json:"@odata.type,omitempty"`
	ResultType      *ThreatAssessmentResultResultType `json:"resultType,omitempty"`
}
