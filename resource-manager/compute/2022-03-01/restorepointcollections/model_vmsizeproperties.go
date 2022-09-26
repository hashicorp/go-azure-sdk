package restorepointcollections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VmSizeProperties struct {
	VCPUsAvailable *int64 `json:"vCPUsAvailable,omitempty"`
	VCPUsPerCore   *int64 `json:"vCPUsPerCore,omitempty"`
}
