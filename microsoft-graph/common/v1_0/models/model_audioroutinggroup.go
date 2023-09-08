package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudioRoutingGroup struct {
	Id          *string                       `json:"id,omitempty"`
	ODataType   *string                       `json:"@odata.type,omitempty"`
	Receivers   *[]string                     `json:"receivers,omitempty"`
	RoutingMode *AudioRoutingGroupRoutingMode `json:"routingMode,omitempty"`
	Sources     *[]string                     `json:"sources,omitempty"`
}
