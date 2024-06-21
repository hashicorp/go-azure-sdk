package communitygalleryimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImageProperties struct {
	Architecture        *Architecture                    `json:"architecture,omitempty"`
	ArtifactTags        *map[string]string               `json:"artifactTags,omitempty"`
	Disallowed          *Disallowed                      `json:"disallowed,omitempty"`
	Disclaimer          *CommunityGalleryDisclaimer      `json:"disclaimer,omitempty"`
	EndOfLifeDate       *string                          `json:"endOfLifeDate,omitempty"`
	Eula                *string                          `json:"eula,omitempty"`
	Features            *[]GalleryImageFeature           `json:"features,omitempty"`
	HyperVGeneration    *HyperVGeneration                `json:"hyperVGeneration,omitempty"`
	Identifier          CommunityGalleryImageIdentifier  `json:"identifier"`
	OsState             OperatingSystemStateTypes        `json:"osState"`
	OsType              OperatingSystemTypes             `json:"osType"`
	PrivacyStatementUri *string                          `json:"privacyStatementUri,omitempty"`
	PurchasePlan        *ImagePurchasePlan               `json:"purchasePlan,omitempty"`
	Recommended         *RecommendedMachineConfiguration `json:"recommended,omitempty"`
}
