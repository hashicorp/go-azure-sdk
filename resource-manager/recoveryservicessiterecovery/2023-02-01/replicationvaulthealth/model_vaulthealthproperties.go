package replicationvaulthealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultHealthProperties struct {
	ContainersHealth     *ResourceHealthSummary `json:"containersHealth,omitempty"`
	FabricsHealth        *ResourceHealthSummary `json:"fabricsHealth,omitempty"`
	ProtectedItemsHealth *ResourceHealthSummary `json:"protectedItemsHealth,omitempty"`
	VaultErrors          *[]HealthError         `json:"vaultErrors,omitempty"`
}
