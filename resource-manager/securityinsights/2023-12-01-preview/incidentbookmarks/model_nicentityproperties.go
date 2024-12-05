package incidentbookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NicEntityProperties struct {
	AdditionalData    *interface{} `json:"additionalData,omitempty"`
	FriendlyName      *string      `json:"friendlyName,omitempty"`
	IPAddressEntityId *string      `json:"ipAddressEntityId,omitempty"`
	MacAddress        *string      `json:"macAddress,omitempty"`
	Vlans             *[]string    `json:"vlans,omitempty"`
}
