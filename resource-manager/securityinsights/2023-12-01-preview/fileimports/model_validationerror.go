package fileimports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationError struct {
	ErrorMessages *[]string `json:"errorMessages,omitempty"`
	RecordIndex   *int64    `json:"recordIndex,omitempty"`
}
