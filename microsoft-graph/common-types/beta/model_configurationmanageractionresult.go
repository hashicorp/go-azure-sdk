package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceActionResult = ConfigurationManagerActionResult{}

type ConfigurationManagerActionResult struct {
	// Delivery state of Configuration Manager device action
	ActionDeliveryStatus *ConfigurationManagerActionDeliveryStatus `json:"actionDeliveryStatus,omitempty"`

	// Error code of Configuration Manager action from client
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// Fields inherited from DeviceActionResult

	// Action name
	ActionName nullable.Type[string] `json:"actionName,omitempty"`

	ActionState *ActionState `json:"actionState,omitempty"`

	// Time the action state was last updated
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time the action was initiated
	StartDateTime *string `json:"startDateTime,omitempty"`
}

func (s ConfigurationManagerActionResult) DeviceActionResult() BaseDeviceActionResultImpl {
	return BaseDeviceActionResultImpl{
		ActionName:          s.ActionName,
		ActionState:         s.ActionState,
		LastUpdatedDateTime: s.LastUpdatedDateTime,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		StartDateTime:       s.StartDateTime,
	}
}

var _ json.Marshaler = ConfigurationManagerActionResult{}

func (s ConfigurationManagerActionResult) MarshalJSON() ([]byte, error) {
	type wrapper ConfigurationManagerActionResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConfigurationManagerActionResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConfigurationManagerActionResult: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.configurationManagerActionResult"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConfigurationManagerActionResult: %+v", err)
	}

	return encoded, nil
}
