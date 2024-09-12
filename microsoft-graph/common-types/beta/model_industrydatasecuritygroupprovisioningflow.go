package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IndustryDataProvisioningFlow = IndustryDataSecurityGroupProvisioningFlow{}

type IndustryDataSecurityGroupProvisioningFlow struct {
	CreationOptions *IndustryDataSecurityGroupCreationOptions `json:"creationOptions,omitempty"`

	// Fields inherited from IndustryDataProvisioningFlow

	// The date and time when the provisioning flow was created. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The date and time when the provisioning flow was most recently changed. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The state of the activity from creation through to ready to do work. The possible values are: notReady, ready,
	// failed, disabled, expired, unknownFutureValue.
	ReadinessStatus *IndustryDataReadinessStatus `json:"readinessStatus,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s IndustryDataSecurityGroupProvisioningFlow) IndustryDataProvisioningFlow() BaseIndustryDataProvisioningFlowImpl {
	return BaseIndustryDataProvisioningFlowImpl{
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ReadinessStatus:      s.ReadinessStatus,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s IndustryDataSecurityGroupProvisioningFlow) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataSecurityGroupProvisioningFlow{}

func (s IndustryDataSecurityGroupProvisioningFlow) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataSecurityGroupProvisioningFlow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataSecurityGroupProvisioningFlow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataSecurityGroupProvisioningFlow: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.industryData.securityGroupProvisioningFlow"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataSecurityGroupProvisioningFlow: %+v", err)
	}

	return encoded, nil
}
