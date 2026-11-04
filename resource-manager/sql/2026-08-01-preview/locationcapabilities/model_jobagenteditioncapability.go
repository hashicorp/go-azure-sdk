package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobAgentEditionCapability struct {
	Name                            *string                                    `json:"name,omitempty"`
	Reason                          *string                                    `json:"reason,omitempty"`
	Status                          *CapabilityStatus                          `json:"status,omitempty"`
	SupportedServiceLevelObjectives *[]JobAgentServiceLevelObjectiveCapability `json:"supportedServiceLevelObjectives,omitempty"`
}
