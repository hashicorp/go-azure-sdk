package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchError struct {
	Code    *string             `json:"code,omitempty"`
	Message *ErrorMessage       `json:"message,omitempty"`
	Values  *[]BatchErrorDetail `json:"values,omitempty"`
}
