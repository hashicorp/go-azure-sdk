package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersistentBrowserSessionControl struct {
	IsEnabled *bool                                `json:"isEnabled,omitempty"`
	Mode      *PersistentBrowserSessionControlMode `json:"mode,omitempty"`
	ODataType *string                              `json:"@odata.type,omitempty"`
}
