package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DateTimeColumn struct {
	DisplayAs *string `json:"displayAs,omitempty"`
	Format    *string `json:"format,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
