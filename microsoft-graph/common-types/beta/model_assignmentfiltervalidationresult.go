package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterValidationResult struct {
	IsValidRule *bool   `json:"isValidRule,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
