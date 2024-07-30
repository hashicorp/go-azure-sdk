package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientHealthState struct {
	ErrorCode        *int64                                      `json:"errorCode,omitempty"`
	LastSyncDateTime *string                                     `json:"lastSyncDateTime,omitempty"`
	ODataType        *string                                     `json:"@odata.type,omitempty"`
	State            *ConfigurationManagerClientHealthStateState `json:"state,omitempty"`
}
