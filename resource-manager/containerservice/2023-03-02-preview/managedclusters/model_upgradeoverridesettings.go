package managedclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpgradeOverrideSettings struct {
	ControlPlaneOverrides *[]ControlPlaneUpgradeOverride `json:"controlPlaneOverrides,omitempty"`
	Until                 *string                        `json:"until,omitempty"`
}
