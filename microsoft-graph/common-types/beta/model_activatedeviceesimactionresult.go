package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateDeviceEsimActionResult struct {
	ActionName          *string                                    `json:"actionName,omitempty"`
	ActionState         *ActivateDeviceEsimActionResultActionState `json:"actionState,omitempty"`
	CarrierUrl          *string                                    `json:"carrierUrl,omitempty"`
	LastUpdatedDateTime *string                                    `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                                    `json:"@odata.type,omitempty"`
	StartDateTime       *string                                    `json:"startDateTime,omitempty"`
}
