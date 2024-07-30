package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluateLabelJobResultGroup struct {
	Automatic   *EvaluateLabelJobResult `json:"automatic,omitempty"`
	ODataType   *string                 `json:"@odata.type,omitempty"`
	Recommended *EvaluateLabelJobResult `json:"recommended,omitempty"`
}
