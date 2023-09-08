package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIsolateDeviceResponseAction struct {
	Identifier    *SecurityIsolateDeviceResponseActionIdentifier    `json:"identifier,omitempty"`
	IsolationType *SecurityIsolateDeviceResponseActionIsolationType `json:"isolationType,omitempty"`
	ODataType     *string                                           `json:"@odata.type,omitempty"`
}
