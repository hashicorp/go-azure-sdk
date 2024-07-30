package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSet struct {
	AllowedContentTypes         *[]ContentTypeInfo    `json:"allowedContentTypes,omitempty"`
	DefaultContents             *[]DocumentSetContent `json:"defaultContents,omitempty"`
	ODataType                   *string               `json:"@odata.type,omitempty"`
	PropagateWelcomePageChanges *bool                 `json:"propagateWelcomePageChanges,omitempty"`
	SharedColumns               *[]ColumnDefinition   `json:"sharedColumns,omitempty"`
	ShouldPrefixNameToFile      *bool                 `json:"shouldPrefixNameToFile,omitempty"`
	WelcomePageColumns          *[]ColumnDefinition   `json:"welcomePageColumns,omitempty"`
	WelcomePageUrl              *string               `json:"welcomePageUrl,omitempty"`
}
