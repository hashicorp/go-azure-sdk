package changedatacapturecapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureCapabilityPropertiesForPatchUpdate struct {
	AutoUpgradeExtension *bool                          `json:"autoUpgradeExtension,omitempty"`
	DesiredState         *ChangeDataCaptureDesiredState `json:"desiredState,omitempty"`
}
