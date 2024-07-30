package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthoredNote struct {
	Author          *Identity `json:"author,omitempty"`
	Content         *ItemBody `json:"content,omitempty"`
	CreatedDateTime *string   `json:"createdDateTime,omitempty"`
	Id              *string   `json:"id,omitempty"`
	ODataType       *string   `json:"@odata.type,omitempty"`
}
