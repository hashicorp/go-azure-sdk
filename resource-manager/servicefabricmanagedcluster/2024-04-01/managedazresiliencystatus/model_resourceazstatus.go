package managedazresiliencystatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceAzStatus struct {
	Details         *string `json:"details,omitempty"`
	IsZoneResilient *bool   `json:"isZoneResilient,omitempty"`
	ResourceName    *string `json:"resourceName,omitempty"`
	ResourceType    *string `json:"resourceType,omitempty"`
}
