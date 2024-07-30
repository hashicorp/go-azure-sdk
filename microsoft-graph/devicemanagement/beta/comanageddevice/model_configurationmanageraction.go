package comanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerAction struct {
	Action    *ConfigurationManagerActionAction `json:"action,omitempty"`
	ODataType *string                           `json:"@odata.type,omitempty"`
}
