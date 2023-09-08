package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRoute struct {
	Final       *IdentitySet          `json:"final,omitempty"`
	ODataType   *string               `json:"@odata.type,omitempty"`
	Original    *IdentitySet          `json:"original,omitempty"`
	RoutingType *CallRouteRoutingType `json:"routingType,omitempty"`
}
