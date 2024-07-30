package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaArtifact struct {
	BoardSupportPackageVersion *string `json:"boardSupportPackageVersion,omitempty"`
	Description                *string `json:"description,omitempty"`
	DeviceModel                *string `json:"deviceModel,omitempty"`
	Id                         *string `json:"id,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
	OsVersion                  *string `json:"osVersion,omitempty"`
	PatchVersion               *string `json:"patchVersion,omitempty"`
	ReleaseNotesUrl            *string `json:"releaseNotesUrl,omitempty"`
}
