package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ControlPlaneNodePatchConfiguration struct {
	AdministratorConfiguration *AdministratorConfigurationPatch `json:"administratorConfiguration,omitempty"`
	Count                      *int64                           `json:"count,omitempty"`
}
