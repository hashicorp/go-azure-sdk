package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluateRequest struct {
	ODataType *string                                  `json:"@odata.type,omitempty"`
	OrderBy   *[]string                                `json:"orderBy,omitempty"`
	Platform  *AssignmentFilterEvaluateRequestPlatform `json:"platform,omitempty"`
	Rule      *string                                  `json:"rule,omitempty"`
	Search    *string                                  `json:"search,omitempty"`
	Skip      *int64                                   `json:"skip,omitempty"`
	Top       *int64                                   `json:"top,omitempty"`
}
