package dashboards

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DashboardPropertiesWithProvisioningState struct {
	Lenses            *[]DashboardLens           `json:"lenses,omitempty"`
	Metadata          *interface{}               `json:"metadata,omitempty"`
	ProvisioningState *ResourceProvisioningState `json:"provisioningState,omitempty"`
}
