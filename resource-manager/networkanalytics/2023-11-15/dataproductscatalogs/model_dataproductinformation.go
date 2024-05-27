package dataproductscatalogs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProductInformation struct {
	DataProductName     string               `json:"dataProductName"`
	DataProductVersions []DataProductVersion `json:"dataProductVersions"`
	Description         string               `json:"description"`
}
