package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConfiguredPeripheral struct {
	IsOptional *bool               `json:"isOptional,omitempty"`
	ODataType  *string             `json:"@odata.type,omitempty"`
	Peripheral *TeamworkPeripheral `json:"peripheral,omitempty"`
}
