package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningProvider struct {
	DisplayName                   *string                   `json:"displayName,omitempty"`
	Id                            *string                   `json:"id,omitempty"`
	IsCourseActivitySyncEnabled   *bool                     `json:"isCourseActivitySyncEnabled,omitempty"`
	LearningContents              *[]LearningContent        `json:"learningContents,omitempty"`
	LearningCourseActivities      *[]LearningCourseActivity `json:"learningCourseActivities,omitempty"`
	LoginWebUrl                   *string                   `json:"loginWebUrl,omitempty"`
	LongLogoWebUrlForDarkTheme    *string                   `json:"longLogoWebUrlForDarkTheme,omitempty"`
	LongLogoWebUrlForLightTheme   *string                   `json:"longLogoWebUrlForLightTheme,omitempty"`
	ODataType                     *string                   `json:"@odata.type,omitempty"`
	SquareLogoWebUrlForDarkTheme  *string                   `json:"squareLogoWebUrlForDarkTheme,omitempty"`
	SquareLogoWebUrlForLightTheme *string                   `json:"squareLogoWebUrlForLightTheme,omitempty"`
}
