package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentState struct {
	EffectiveValue *WindowsUpdatesDeploymentStateEffectiveValue `json:"effectiveValue,omitempty"`
	ODataType      *string                                      `json:"@odata.type,omitempty"`
	Reasons        *[]WindowsUpdatesDeploymentStateReason       `json:"reasons,omitempty"`
	RequestedValue *WindowsUpdatesDeploymentStateRequestedValue `json:"requestedValue,omitempty"`
}
