package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertComment struct {
	Comment              *string `json:"comment,omitempty"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	CreatedDateTime      *string `json:"createdDateTime,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
