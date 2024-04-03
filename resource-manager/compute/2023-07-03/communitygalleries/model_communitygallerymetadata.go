package communitygalleries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryMetadata struct {
	Eula                *string  `json:"eula,omitempty"`
	PrivacyStatementUri *string  `json:"privacyStatementUri,omitempty"`
	PublicNames         []string `json:"publicNames"`
	PublisherContact    string   `json:"publisherContact"`
	PublisherUri        *string  `json:"publisherUri,omitempty"`
}
