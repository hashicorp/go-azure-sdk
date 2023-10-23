package sku

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationInfo struct {
	Capabilities *[]Capability `json:"capabilities,omitempty"`
	Location     *string       `json:"location,omitempty"`
}
