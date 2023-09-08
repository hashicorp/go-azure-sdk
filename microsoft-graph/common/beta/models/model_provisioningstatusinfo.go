package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStatusInfo struct {
	ErrorInformation *ProvisioningErrorInfo        `json:"errorInformation,omitempty"`
	ODataType        *string                       `json:"@odata.type,omitempty"`
	Status           *ProvisioningStatusInfoStatus `json:"status,omitempty"`
}
