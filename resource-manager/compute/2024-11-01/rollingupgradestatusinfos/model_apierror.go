package rollingupgradestatusinfos

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiError struct {
	Code       *string         `json:"code,omitempty"`
	Details    *[]ApiErrorBase `json:"details,omitempty"`
	Innererror *InnerError     `json:"innererror,omitempty"`
	Message    *string         `json:"message,omitempty"`
	Target     *string         `json:"target,omitempty"`
}
