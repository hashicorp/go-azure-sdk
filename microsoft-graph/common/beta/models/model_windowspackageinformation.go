package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPackageInformation struct {
	ApplicableArchitecture          *WindowsPackageInformationApplicableArchitecture `json:"applicableArchitecture,omitempty"`
	DisplayName                     *string                                          `json:"displayName,omitempty"`
	IdentityName                    *string                                          `json:"identityName,omitempty"`
	IdentityPublisher               *string                                          `json:"identityPublisher,omitempty"`
	IdentityResourceIdentifier      *string                                          `json:"identityResourceIdentifier,omitempty"`
	IdentityVersion                 *string                                          `json:"identityVersion,omitempty"`
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem                   `json:"minimumSupportedOperatingSystem,omitempty"`
	ODataType                       *string                                          `json:"@odata.type,omitempty"`
}
