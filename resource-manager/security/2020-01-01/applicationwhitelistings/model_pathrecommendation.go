package applicationwhitelistings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PathRecommendation struct {
	Action              *Action               `json:"action,omitempty"`
	Common              *bool                 `json:"common,omitempty"`
	ConfigurationStatus *ConfigurationStatus  `json:"configurationStatus,omitempty"`
	FileType            *FileType             `json:"fileType,omitempty"`
	Path                *string               `json:"path,omitempty"`
	PublisherInfo       *PublisherInfo        `json:"publisherInfo,omitempty"`
	Type                *Type                 `json:"type,omitempty"`
	UserSids            *[]string             `json:"userSids,omitempty"`
	Usernames           *[]UserRecommendation `json:"usernames,omitempty"`
}
