package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProvisioningError struct {
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	IsResolved      *bool   `json:"isResolved,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	ServiceInstance *string `json:"serviceInstance,omitempty"`
}
