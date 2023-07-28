package dataversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntellectualProperty struct {
	ProtectionLevel *ProtectionLevel `json:"protectionLevel,omitempty"`
	Publisher       string           `json:"publisher"`
}
