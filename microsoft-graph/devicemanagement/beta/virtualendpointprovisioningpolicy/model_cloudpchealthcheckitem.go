package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCHealthCheckItem struct {
	AdditionalDetails       *string                       `json:"additionalDetails,omitempty"`
	DisplayName             *string                       `json:"displayName,omitempty"`
	LastHealthCheckDateTime *string                       `json:"lastHealthCheckDateTime,omitempty"`
	ODataType               *string                       `json:"@odata.type,omitempty"`
	Result                  *CloudPCHealthCheckItemResult `json:"result,omitempty"`
}
