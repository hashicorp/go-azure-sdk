package integrationserviceenvironmentnetworkhealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendedErrorInfo struct {
	Code       ErrorResponseCode    `json:"code"`
	Details    *[]ExtendedErrorInfo `json:"details,omitempty"`
	InnerError *interface{}         `json:"innerError,omitempty"`
	Message    string               `json:"message"`
}
