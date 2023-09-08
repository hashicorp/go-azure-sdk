package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssuePost struct {
	CreatedDateTime *string                         `json:"createdDateTime,omitempty"`
	Description     *ItemBody                       `json:"description,omitempty"`
	ODataType       *string                         `json:"@odata.type,omitempty"`
	PostType        *ServiceHealthIssuePostPostType `json:"postType,omitempty"`
}
