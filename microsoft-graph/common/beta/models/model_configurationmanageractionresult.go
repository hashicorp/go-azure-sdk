package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionResult struct {
	ActionDeliveryStatus *ConfigurationManagerActionResultActionDeliveryStatus `json:"actionDeliveryStatus,omitempty"`
	ActionName           *string                                               `json:"actionName,omitempty"`
	ActionState          *ConfigurationManagerActionResultActionState          `json:"actionState,omitempty"`
	ErrorCode            *int64                                                `json:"errorCode,omitempty"`
	LastUpdatedDateTime  *string                                               `json:"lastUpdatedDateTime,omitempty"`
	ODataType            *string                                               `json:"@odata.type,omitempty"`
	StartDateTime        *string                                               `json:"startDateTime,omitempty"`
}
