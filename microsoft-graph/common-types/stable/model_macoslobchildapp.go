package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSLobChildApp struct {
	BuildNumber   *string `json:"buildNumber,omitempty"`
	BundleId      *string `json:"bundleId,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	VersionNumber *string `json:"versionNumber,omitempty"`
}
