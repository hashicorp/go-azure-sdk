package communitygalleryimages

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImageProperties struct {
	Disallowed       *Disallowed                      `json:"disallowed,omitempty"`
	EndOfLifeDate    *string                          `json:"endOfLifeDate,omitempty"`
	Features         *[]GalleryImageFeature           `json:"features,omitempty"`
	HyperVGeneration *HyperVGeneration                `json:"hyperVGeneration,omitempty"`
	Identifier       GalleryImageIdentifier           `json:"identifier"`
	OsState          OperatingSystemStateTypes        `json:"osState"`
	OsType           OperatingSystemTypes             `json:"osType"`
	PurchasePlan     *ImagePurchasePlan               `json:"purchasePlan,omitempty"`
	Recommended      *RecommendedMachineConfiguration `json:"recommended,omitempty"`
}

func (o *CommunityGalleryImageProperties) GetEndOfLifeDateAsTime() (*time.Time, error) {
	if o.EndOfLifeDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndOfLifeDate, "2006-01-02T15:04:05Z07:00")
}

func (o *CommunityGalleryImageProperties) SetEndOfLifeDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndOfLifeDate = &formatted
}
