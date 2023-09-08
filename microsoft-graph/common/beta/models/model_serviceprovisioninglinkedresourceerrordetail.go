package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProvisioningLinkedResourceErrorDetail struct {
	Code         *string `json:"code,omitempty"`
	Details      *string `json:"details,omitempty"`
	Message      *string `json:"message,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	PropertyName *string `json:"propertyName,omitempty"`
	Target       *string `json:"target,omitempty"`
}
