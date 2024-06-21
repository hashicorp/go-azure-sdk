package communitygalleryimages

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
