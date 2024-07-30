package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSKernelExtension struct {
	BundleId       *string `json:"bundleId,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	TeamIdentifier *string `json:"teamIdentifier,omitempty"`
}
