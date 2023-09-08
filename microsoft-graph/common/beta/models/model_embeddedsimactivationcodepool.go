package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMActivationCodePool struct {
	ActivationCodeCount *int64                                     `json:"activationCodeCount,omitempty"`
	ActivationCodes     *[]EmbeddedSIMActivationCode               `json:"activationCodes,omitempty"`
	Assignments         *[]EmbeddedSIMActivationCodePoolAssignment `json:"assignments,omitempty"`
	CreatedDateTime     *string                                    `json:"createdDateTime,omitempty"`
	DeviceStates        *[]EmbeddedSIMDeviceState                  `json:"deviceStates,omitempty"`
	DisplayName         *string                                    `json:"displayName,omitempty"`
	Id                  *string                                    `json:"id,omitempty"`
	ModifiedDateTime    *string                                    `json:"modifiedDateTime,omitempty"`
	ODataType           *string                                    `json:"@odata.type,omitempty"`
}
