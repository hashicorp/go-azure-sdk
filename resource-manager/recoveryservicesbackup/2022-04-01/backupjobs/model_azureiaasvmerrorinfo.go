package backupjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureIaaSVMErrorInfo struct {
	ErrorCode       *int64    `json:"errorCode,omitempty"`
	ErrorString     *string   `json:"errorString,omitempty"`
	ErrorTitle      *string   `json:"errorTitle,omitempty"`
	Recommendations *[]string `json:"recommendations,omitempty"`
}
