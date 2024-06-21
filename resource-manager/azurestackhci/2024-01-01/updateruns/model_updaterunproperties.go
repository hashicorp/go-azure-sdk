package updateruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRunProperties struct {
	Duration          *string                   `json:"duration,omitempty"`
	LastUpdatedTime   *string                   `json:"lastUpdatedTime,omitempty"`
	Progress          *Step                     `json:"progress,omitempty"`
	ProvisioningState *ProvisioningState        `json:"provisioningState,omitempty"`
	State             *UpdateRunPropertiesState `json:"state,omitempty"`
	TimeStarted       *string                   `json:"timeStarted,omitempty"`
}
