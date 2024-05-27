package dataproducts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProductNetworkAcls struct {
	AllowedQueryIPRangeList []string             `json:"allowedQueryIpRangeList"`
	DefaultAction           DefaultAction        `json:"defaultAction"`
	IPRules                 []IPRules            `json:"ipRules"`
	VirtualNetworkRule      []VirtualNetworkRule `json:"virtualNetworkRule"`
}
