package appattachpackageinfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportPackageInfoRequest struct {
	PackageArchitecture *AppAttachPackageArchitectures `json:"packageArchitecture,omitempty"`
	Path                *string                        `json:"path,omitempty"`
}
