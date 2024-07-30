package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRemediationSummary struct {
	ODataType             *string `json:"@odata.type,omitempty"`
	RemediatedDeviceCount *int64  `json:"remediatedDeviceCount,omitempty"`
	ScriptCount           *int64  `json:"scriptCount,omitempty"`
}
