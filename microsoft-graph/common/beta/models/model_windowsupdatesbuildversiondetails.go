package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesBuildVersionDetails struct {
	BuildNumber         *int64  `json:"buildNumber,omitempty"`
	MajorVersion        *int64  `json:"majorVersion,omitempty"`
	MinorVersion        *int64  `json:"minorVersion,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	UpdateBuildRevision *int64  `json:"updateBuildRevision,omitempty"`
}
