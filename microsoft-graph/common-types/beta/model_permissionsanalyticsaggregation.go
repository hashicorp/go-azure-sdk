package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticsAggregation struct {
	Aws       *PermissionsAnalytics `json:"aws,omitempty"`
	Azure     *PermissionsAnalytics `json:"azure,omitempty"`
	Gcp       *PermissionsAnalytics `json:"gcp,omitempty"`
	Id        *string               `json:"id,omitempty"`
	ODataType *string               `json:"@odata.type,omitempty"`
}
