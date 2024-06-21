package galleryimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GalleryImageProperties struct {
	Description         *string                          `json:"description,omitempty"`
	Disallowed          *Disallowed                      `json:"disallowed,omitempty"`
	EndOfLifeDate       *string                          `json:"endOfLifeDate,omitempty"`
	Eula                *string                          `json:"eula,omitempty"`
	Features            *[]GalleryImageFeature           `json:"features,omitempty"`
	HyperVGeneration    *HyperVGeneration                `json:"hyperVGeneration,omitempty"`
	Identifier          GalleryImageIdentifier           `json:"identifier"`
	OsState             OperatingSystemStateTypes        `json:"osState"`
	OsType              OperatingSystemTypes             `json:"osType"`
	PrivacyStatementUri *string                          `json:"privacyStatementUri,omitempty"`
	ProvisioningState   *ProvisioningState               `json:"provisioningState,omitempty"`
	PurchasePlan        *ImagePurchasePlan               `json:"purchasePlan,omitempty"`
	Recommended         *RecommendedMachineConfiguration `json:"recommended,omitempty"`
	ReleaseNoteUri      *string                          `json:"releaseNoteUri,omitempty"`
}
