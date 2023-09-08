package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Note struct {
	Attachments                   *[]Attachment                        `json:"attachments,omitempty"`
	Body                          *ItemBody                            `json:"body,omitempty"`
	Categories                    *[]string                            `json:"categories,omitempty"`
	ChangeKey                     *string                              `json:"changeKey,omitempty"`
	CreatedDateTime               *string                              `json:"createdDateTime,omitempty"`
	Extensions                    *[]Extension                         `json:"extensions,omitempty"`
	HasAttachments                *bool                                `json:"hasAttachments,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	IsDeleted                     *bool                                `json:"isDeleted,omitempty"`
	LastModifiedDateTime          *string                              `json:"lastModifiedDateTime,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	Subject                       *string                              `json:"subject,omitempty"`
}
