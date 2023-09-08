package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDomainJoinConnector struct {
	DisplayName            *string                                   `json:"displayName,omitempty"`
	Id                     *string                                   `json:"id,omitempty"`
	LastConnectionDateTime *string                                   `json:"lastConnectionDateTime,omitempty"`
	ODataType              *string                                   `json:"@odata.type,omitempty"`
	State                  *DeviceManagementDomainJoinConnectorState `json:"state,omitempty"`
	Version                *string                                   `json:"version,omitempty"`
}
