package exports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportSuspensionContext struct {
	SuspensionCode   *string `json:"suspensionCode,omitempty"`
	SuspensionReason *string `json:"suspensionReason,omitempty"`
	SuspensionTime   *string `json:"suspensionTime,omitempty"`
}
