package apirevision

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiRevisionContract struct {
	ApiId           *string `json:"apiId,omitempty"`
	ApiRevision     *string `json:"apiRevision,omitempty"`
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	Description     *string `json:"description,omitempty"`
	IsCurrent       *bool   `json:"isCurrent,omitempty"`
	IsOnline        *bool   `json:"isOnline,omitempty"`
	PrivateUrl      *string `json:"privateUrl,omitempty"`
	UpdatedDateTime *string `json:"updatedDateTime,omitempty"`
}
