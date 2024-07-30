package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessUser struct {
	DisplayName        *string                       `json:"displayName,omitempty"`
	LastAccessDateTime *string                       `json:"lastAccessDateTime,omitempty"`
	ODataType          *string                       `json:"@odata.type,omitempty"`
	TrafficType        *NetworkaccessUserTrafficType `json:"trafficType,omitempty"`
	UserId             *string                       `json:"userId,omitempty"`
	UserPrincipalName  *string                       `json:"userPrincipalName,omitempty"`
	UserType           *NetworkaccessUserUserType    `json:"userType,omitempty"`
}
