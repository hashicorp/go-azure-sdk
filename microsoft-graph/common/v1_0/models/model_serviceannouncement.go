package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceAnnouncement struct {
	HealthOverviews *[]ServiceHealth        `json:"healthOverviews,omitempty"`
	Id              *string                 `json:"id,omitempty"`
	Issues          *[]ServiceHealthIssue   `json:"issues,omitempty"`
	Messages        *[]ServiceUpdateMessage `json:"messages,omitempty"`
	ODataType       *string                 `json:"@odata.type,omitempty"`
}
