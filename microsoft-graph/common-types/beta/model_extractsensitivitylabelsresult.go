package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtractSensitivityLabelsResult struct {
	Labels    *[]SensitivityLabelAssignment `json:"labels,omitempty"`
	ODataType *string                       `json:"@odata.type,omitempty"`
}
