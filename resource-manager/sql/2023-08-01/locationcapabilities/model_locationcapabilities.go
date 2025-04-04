package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationCapabilities struct {
	Name                             *string                             `json:"name,omitempty"`
	Reason                           *string                             `json:"reason,omitempty"`
	Status                           *CapabilityStatus                   `json:"status,omitempty"`
	SupportedJobAgentVersions        *[]JobAgentVersionCapability        `json:"supportedJobAgentVersions,omitempty"`
	SupportedManagedInstanceVersions *[]ManagedInstanceVersionCapability `json:"supportedManagedInstanceVersions,omitempty"`
	SupportedServerVersions          *[]ServerVersionCapability          `json:"supportedServerVersions,omitempty"`
}
