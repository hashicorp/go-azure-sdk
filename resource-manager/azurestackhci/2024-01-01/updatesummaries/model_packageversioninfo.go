package updatesummaries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PackageVersionInfo struct {
	LastUpdated *string `json:"lastUpdated,omitempty"`
	PackageType *string `json:"packageType,omitempty"`
	Version     *string `json:"version,omitempty"`
}
