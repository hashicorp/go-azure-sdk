package securitysolutionsreferencedata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySolutionsReferenceDataProperties struct {
	AlertVendorName      string         `json:"alertVendorName"`
	PackageInfoURL       string         `json:"packageInfoUrl"`
	ProductName          string         `json:"productName"`
	Publisher            string         `json:"publisher"`
	PublisherDisplayName string         `json:"publisherDisplayName"`
	SecurityFamily       SecurityFamily `json:"securityFamily"`
	Template             string         `json:"template"`
}
