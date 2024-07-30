package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSIncludedApp struct {
	BundleId      *string `json:"bundleId,omitempty"`
	BundleVersion *string `json:"bundleVersion,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
