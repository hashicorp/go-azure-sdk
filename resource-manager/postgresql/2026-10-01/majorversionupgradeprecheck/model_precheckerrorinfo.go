package majorversionupgradeprecheck

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrecheckErrorInfo struct {
	ErrorCode    *int64  `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
}
