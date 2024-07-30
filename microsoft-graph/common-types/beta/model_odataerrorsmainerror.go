package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ODataErrorsMainError struct {
	Code       *string                    `json:"code,omitempty"`
	Details    *[]ODataErrorsErrorDetails `json:"details,omitempty"`
	InnerError *InnerError                `json:"innerError,omitempty"`
	Message    *string                    `json:"message,omitempty"`
	Target     *string                    `json:"target,omitempty"`
}
