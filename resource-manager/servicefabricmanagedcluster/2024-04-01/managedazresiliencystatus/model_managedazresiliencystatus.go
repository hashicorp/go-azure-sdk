package managedazresiliencystatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAzResiliencyStatus struct {
	BaseResourceStatus     *[]ResourceAzStatus `json:"baseResourceStatus,omitempty"`
	IsClusterZoneResilient *bool               `json:"isClusterZoneResilient,omitempty"`
}
