package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeActionResult struct {
	DeviceScopeAction *string                        `json:"deviceScopeAction,omitempty"`
	DeviceScopeId     *string                        `json:"deviceScopeId,omitempty"`
	FailedMessage     *string                        `json:"failedMessage,omitempty"`
	ODataType         *string                        `json:"@odata.type,omitempty"`
	Status            *DeviceScopeActionResultStatus `json:"status,omitempty"`
}
