package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceAnnouncementBase struct {
	Details              *[]KeyValuePair `json:"details,omitempty"`
	EndDateTime          *string         `json:"endDateTime,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	LastModifiedDateTime *string         `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string         `json:"@odata.type,omitempty"`
	StartDateTime        *string         `json:"startDateTime,omitempty"`
	Title                *string         `json:"title,omitempty"`
}
