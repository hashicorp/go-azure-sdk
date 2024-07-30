package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAppXAppAssignmentSettings struct {
	ODataType        *string `json:"@odata.type,omitempty"`
	UseDeviceContext *bool   `json:"useDeviceContext,omitempty"`
}
