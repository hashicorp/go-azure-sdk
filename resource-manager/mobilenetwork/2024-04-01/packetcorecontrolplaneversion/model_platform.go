package packetcorecontrolplaneversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Platform struct {
	HaUpgradesAvailable            *[]string           `json:"haUpgradesAvailable,omitempty"`
	MaximumPlatformSoftwareVersion *string             `json:"maximumPlatformSoftwareVersion,omitempty"`
	MinimumPlatformSoftwareVersion *string             `json:"minimumPlatformSoftwareVersion,omitempty"`
	ObsoleteVersion                *ObsoleteVersion    `json:"obsoleteVersion,omitempty"`
	PlatformType                   *PlatformType       `json:"platformType,omitempty"`
	RecommendedVersion             *RecommendedVersion `json:"recommendedVersion,omitempty"`
	VersionState                   *VersionState       `json:"versionState,omitempty"`
}
