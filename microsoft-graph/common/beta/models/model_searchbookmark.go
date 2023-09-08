package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchBookmark struct {
	AvailabilityEndDateTime   *string                    `json:"availabilityEndDateTime,omitempty"`
	AvailabilityStartDateTime *string                    `json:"availabilityStartDateTime,omitempty"`
	Categories                *[]string                  `json:"categories,omitempty"`
	Description               *string                    `json:"description,omitempty"`
	DisplayName               *string                    `json:"displayName,omitempty"`
	GroupIds                  *[]string                  `json:"groupIds,omitempty"`
	Id                        *string                    `json:"id,omitempty"`
	IsSuggested               *bool                      `json:"isSuggested,omitempty"`
	Keywords                  *SearchAnswerKeyword       `json:"keywords,omitempty"`
	LanguageTags              *[]string                  `json:"languageTags,omitempty"`
	LastModifiedBy            *SearchIdentitySet         `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime      *string                    `json:"lastModifiedDateTime,omitempty"`
	ODataType                 *string                    `json:"@odata.type,omitempty"`
	Platforms                 *[]SearchBookmarkPlatforms `json:"platforms,omitempty"`
	PowerAppIds               *[]string                  `json:"powerAppIds,omitempty"`
	State                     *SearchBookmarkState       `json:"state,omitempty"`
	TargetedVariations        *[]SearchAnswerVariant     `json:"targetedVariations,omitempty"`
	WebUrl                    *string                    `json:"webUrl,omitempty"`
}
