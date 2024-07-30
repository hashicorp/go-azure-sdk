package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkHardwareDetail struct {
	MacAddresses *[]string `json:"macAddresses,omitempty"`
	Manufacturer *string   `json:"manufacturer,omitempty"`
	Model        *string   `json:"model,omitempty"`
	ODataType    *string   `json:"@odata.type,omitempty"`
	SerialNumber *string   `json:"serialNumber,omitempty"`
	UniqueId     *string   `json:"uniqueId,omitempty"`
}
