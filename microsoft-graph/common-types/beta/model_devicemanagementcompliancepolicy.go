package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementCompliancePolicy{}

type DeviceManagementCompliancePolicy struct {
	// Policy assignments
	Assignments *[]DeviceManagementConfigurationPolicyAssignment `json:"assignments,omitempty"`

	// Policy creation date and time. This property is read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Policy creation source
	CreationSource nullable.Type[string] `json:"creationSource,omitempty"`

	// Policy description
	Description nullable.Type[string] `json:"description,omitempty"`

	// Policy assignment status. This property is read-only.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// Policy last modification date and time. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Policy name
	Name nullable.Type[string] `json:"name,omitempty"`

	// Supported platform types.
	Platforms *DeviceManagementConfigurationPlatforms `json:"platforms,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The list of scheduled action for this rule
	ScheduledActionsForRule *[]DeviceManagementComplianceScheduledActionForRule `json:"scheduledActionsForRule,omitempty"`

	// Number of settings. This property is read-only.
	SettingCount *int64 `json:"settingCount,omitempty"`

	// Policy settings
	Settings *[]DeviceManagementConfigurationSetting `json:"settings,omitempty"`

	// Describes which technology this setting can be deployed with
	Technologies *DeviceManagementConfigurationTechnologies `json:"technologies,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s DeviceManagementCompliancePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementCompliancePolicy{}

func (s DeviceManagementCompliancePolicy) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementCompliancePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementCompliancePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementCompliancePolicy: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "isAssigned")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "settingCount")
	decoded["@odata.type"] = "#microsoft.graph.deviceManagementCompliancePolicy"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementCompliancePolicy: %+v", err)
	}

	return encoded, nil
}
