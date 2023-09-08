package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationError struct {
	Code       *string                   `json:"code,omitempty"`
	Details    *[]ClassifcationErrorBase `json:"details,omitempty"`
	InnerError *ClassificationInnerError `json:"innerError,omitempty"`
	Message    *string                   `json:"message,omitempty"`
	ODataType  *string                   `json:"@odata.type,omitempty"`
	Target     *string                   `json:"target,omitempty"`
}
