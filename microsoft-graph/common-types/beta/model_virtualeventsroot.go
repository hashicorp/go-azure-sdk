package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventsRoot struct {
	Events    *[]VirtualEvent         `json:"events,omitempty"`
	Id        *string                 `json:"id,omitempty"`
	ODataType *string                 `json:"@odata.type,omitempty"`
	Townhalls *[]VirtualEventTownhall `json:"townhalls,omitempty"`
	Webinars  *[]VirtualEventWebinar  `json:"webinars,omitempty"`
}
