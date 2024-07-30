package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCStatusDetails struct {
	AdditionalInformation *[]KeyValuePair `json:"additionalInformation,omitempty"`
	Code                  *string         `json:"code,omitempty"`
	Message               *string         `json:"message,omitempty"`
	ODataType             *string         `json:"@odata.type,omitempty"`
}
