package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DNSZoneResponse struct {
	RequiredZoneNames *[]string             `json:"requiredZoneNames,omitempty"`
	SubResource       *VaultSubResourceType `json:"subResource,omitempty"`
}
