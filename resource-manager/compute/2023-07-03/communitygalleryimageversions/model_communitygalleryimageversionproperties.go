package communitygalleryimageversions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImageVersionProperties struct {
	ArtifactTags      *map[string]string                       `json:"artifactTags,omitempty"`
	Disclaimer        *CommunityGalleryDisclaimer              `json:"disclaimer,omitempty"`
	EndOfLifeDate     *string                                  `json:"endOfLifeDate,omitempty"`
	ExcludeFromLatest *bool                                    `json:"excludeFromLatest,omitempty"`
	PublishedDate     *string                                  `json:"publishedDate,omitempty"`
	StorageProfile    *SharedGalleryImageVersionStorageProfile `json:"storageProfile,omitempty"`
}

func (o *CommunityGalleryImageVersionProperties) GetEndOfLifeDateAsTime() (*time.Time, error) {
	if o.EndOfLifeDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndOfLifeDate, "2006-01-02T15:04:05Z07:00")
}

func (o *CommunityGalleryImageVersionProperties) SetEndOfLifeDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndOfLifeDate = &formatted
}

func (o *CommunityGalleryImageVersionProperties) GetPublishedDateAsTime() (*time.Time, error) {
	if o.PublishedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PublishedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *CommunityGalleryImageVersionProperties) SetPublishedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PublishedDate = &formatted
}
