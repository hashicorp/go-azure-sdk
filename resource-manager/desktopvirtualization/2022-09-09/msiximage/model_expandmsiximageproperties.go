package msiximage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpandMsixImageProperties struct {
	DisplayName           *string                    `json:"displayName,omitempty"`
	ImagePath             *string                    `json:"imagePath,omitempty"`
	IsActive              *bool                      `json:"isActive,omitempty"`
	IsRegularRegistration *bool                      `json:"isRegularRegistration,omitempty"`
	LastUpdated           *string                    `json:"lastUpdated,omitempty"`
	PackageAlias          *string                    `json:"packageAlias,omitempty"`
	PackageApplications   *[]MsixPackageApplications `json:"packageApplications,omitempty"`
	PackageDependencies   *[]MsixPackageDependencies `json:"packageDependencies,omitempty"`
	PackageFamilyName     *string                    `json:"packageFamilyName,omitempty"`
	PackageFullName       *string                    `json:"packageFullName,omitempty"`
	PackageName           *string                    `json:"packageName,omitempty"`
	PackageRelativePath   *string                    `json:"packageRelativePath,omitempty"`
	Version               *string                    `json:"version,omitempty"`
}
