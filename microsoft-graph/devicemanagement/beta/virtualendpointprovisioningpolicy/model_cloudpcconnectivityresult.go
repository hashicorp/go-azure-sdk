package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityResult struct {
	FailedHealthCheckItems *[]CloudPCHealthCheckItem        `json:"failedHealthCheckItems,omitempty"`
	ODataType              *string                          `json:"@odata.type,omitempty"`
	Status                 *CloudPCConnectivityResultStatus `json:"status,omitempty"`
	UpdatedDateTime        *string                          `json:"updatedDateTime,omitempty"`
}
