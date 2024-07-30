package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Connector struct {
	ExternalIp  *string           `json:"externalIp,omitempty"`
	Id          *string           `json:"id,omitempty"`
	MachineName *string           `json:"machineName,omitempty"`
	MemberOf    *[]ConnectorGroup `json:"memberOf,omitempty"`
	ODataType   *string           `json:"@odata.type,omitempty"`
	Status      *ConnectorStatus  `json:"status,omitempty"`
	Version     *string           `json:"version,omitempty"`
}
