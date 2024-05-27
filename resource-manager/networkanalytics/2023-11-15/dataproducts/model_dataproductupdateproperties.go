package dataproducts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProductUpdateProperties struct {
	CurrentMinorVersion *string       `json:"currentMinorVersion,omitempty"`
	Owners              *[]string     `json:"owners,omitempty"`
	PrivateLinksEnabled *ControlState `json:"privateLinksEnabled,omitempty"`
	PurviewAccount      *string       `json:"purviewAccount,omitempty"`
	PurviewCollection   *string       `json:"purviewCollection,omitempty"`
}
