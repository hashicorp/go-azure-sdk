package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolder struct {
	ChildFolderCount              *int64                               `json:"childFolderCount,omitempty"`
	ChildFolders                  *[]MailFolder                        `json:"childFolders,omitempty"`
	DisplayName                   *string                              `json:"displayName,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	IsHidden                      *bool                                `json:"isHidden,omitempty"`
	MessageRules                  *[]MessageRule                       `json:"messageRules,omitempty"`
	Messages                      *[]Message                           `json:"messages,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	ParentFolderId                *string                              `json:"parentFolderId,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	TotalItemCount                *int64                               `json:"totalItemCount,omitempty"`
	UnreadItemCount               *int64                               `json:"unreadItemCount,omitempty"`
}
