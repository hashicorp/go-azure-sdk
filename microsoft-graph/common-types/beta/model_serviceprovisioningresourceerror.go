package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProvisioningResourceError struct {
	CreatedDateTime *string                                   `json:"createdDateTime,omitempty"`
	Errors          *[]ServiceProvisioningResourceErrorDetail `json:"errors,omitempty"`
	IsResolved      *bool                                     `json:"isResolved,omitempty"`
	ODataType       *string                                   `json:"@odata.type,omitempty"`
	ServiceInstance *string                                   `json:"serviceInstance,omitempty"`
}
