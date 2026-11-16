package majorversionupgradeprecheck

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyDetail struct {
	ErrorCode         *int64  `json:"errorCode,omitempty"`
	ErrorMessage      *string `json:"errorMessage,omitempty"`
	Passed            *bool   `json:"passed,omitempty"`
	PolicyDescription *string `json:"policyDescription,omitempty"`
	PolicyName        *string `json:"policyName,omitempty"`
}
