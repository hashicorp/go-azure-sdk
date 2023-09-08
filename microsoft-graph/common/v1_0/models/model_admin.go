package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Admin struct {
	Edge                *Edge                `json:"edge,omitempty"`
	ODataType           *string              `json:"@odata.type,omitempty"`
	ServiceAnnouncement *ServiceAnnouncement `json:"serviceAnnouncement,omitempty"`
	Sharepoint          *Sharepoint          `json:"sharepoint,omitempty"`
}
