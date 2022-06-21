package workspaceskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSkuZoneDetails struct {
	Capabilities *[]SKUCapability `json:"capabilities,omitempty"`
	Name         *[]string        `json:"name,omitempty"`
}
