package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedDeviceMobileAppConfigurationDeviceSummary{}

type ManagedDeviceMobileAppConfigurationDeviceSummary struct {
	// Version of the policy for that overview
	ConfigurationVersion *int64 `json:"configurationVersion,omitempty"`

	// Number of error devices
	ErrorCount *int64 `json:"errorCount,omitempty"`

	// Number of failed devices
	FailedCount *int64 `json:"failedCount,omitempty"`

	// Last update time
	LastUpdateDateTime *string `json:"lastUpdateDateTime,omitempty"`

	// Number of not applicable devices
	NotApplicableCount *int64 `json:"notApplicableCount,omitempty"`

	// Number of pending devices
	PendingCount *int64 `json:"pendingCount,omitempty"`

	// Number of succeeded devices
	SuccessCount *int64 `json:"successCount,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ManagedDeviceMobileAppConfigurationDeviceSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedDeviceMobileAppConfigurationDeviceSummary{}

func (s ManagedDeviceMobileAppConfigurationDeviceSummary) MarshalJSON() ([]byte, error) {
	type wrapper ManagedDeviceMobileAppConfigurationDeviceSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedDeviceMobileAppConfigurationDeviceSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedDeviceMobileAppConfigurationDeviceSummary: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.managedDeviceMobileAppConfigurationDeviceSummary"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedDeviceMobileAppConfigurationDeviceSummary: %+v", err)
	}

	return encoded, nil
}
