package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementExchangeOnPremisesPolicy{}

type DeviceManagementExchangeOnPremisesPolicy struct {
	// The list of device access rules in Exchange. The access rules apply globally to the entire Exchange organization
	AccessRules *[]DeviceManagementExchangeAccessRule `json:"accessRules,omitempty"`

	// The Exchange on premises conditional access settings. On premises conditional access will require devices to be both
	// enrolled and compliant for mail access
	ConditionalAccessSettings *OnPremisesConditionalAccessSettings `json:"conditionalAccessSettings,omitempty"`

	// Access Level in Exchange.
	DefaultAccessLevel *DeviceManagementExchangeAccessLevel `json:"defaultAccessLevel,omitempty"`

	// The list of device classes known to Exchange
	KnownDeviceClasses *[]DeviceManagementExchangeDeviceClass `json:"knownDeviceClasses,omitempty"`

	// Notification text that will be sent to users quarantined by this policy. This is UTF8 encoded byte array HTML.
	NotificationContent nullable.Type[string] `json:"notificationContent,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceManagementExchangeOnPremisesPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementExchangeOnPremisesPolicy{}

func (s DeviceManagementExchangeOnPremisesPolicy) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementExchangeOnPremisesPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementExchangeOnPremisesPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementExchangeOnPremisesPolicy: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.deviceManagementExchangeOnPremisesPolicy"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementExchangeOnPremisesPolicy: %+v", err)
	}

	return encoded, nil
}
