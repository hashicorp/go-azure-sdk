package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrivilegedAccessScheduleInstance = PrivilegedAccessGroupEligibilityScheduleInstance{}

type PrivilegedAccessGroupEligibilityScheduleInstance struct {
	// The identifier of the membership or ownership eligibility relationship to the group. Required. The possible values
	// are: owner, member. Supports $filter (eq).
	AccessId PrivilegedAccessGroupRelationships `json:"accessId"`

	// The identifier of the privilegedAccessGroupEligibilitySchedule from which this instance was created. Required.
	// Supports $filter (eq, ne).
	EligibilityScheduleId nullable.Type[string] `json:"eligibilityScheduleId,omitempty"`

	// References the group that is the scope of the membership or ownership eligibility through PIM for groups. Supports
	// $expand.
	Group *Group `json:"group,omitempty"`

	// The identifier of the group representing the scope of the membership or ownership eligibility through PIM for groups.
	// Required. Supports $filter (eq).
	GroupId nullable.Type[string] `json:"groupId,omitempty"`

	// Indicates whether the assignment is derived from a group assignment. It can further imply whether the calling
	// principal can manage the assignment schedule. Required. The possible values are: direct, group, unknownFutureValue.
	// Supports $filter (eq).
	MemberType PrivilegedAccessGroupMemberType `json:"memberType"`

	// References the principal that's in the scope of the membership or ownership eligibility request through the group
	// that's governed by PIM. Supports $expand.
	Principal *DirectoryObject `json:"principal,omitempty"`

	// The identifier of the principal whose membership or ownership eligibility to the group is managed through PIM for
	// groups. Required. Supports $filter (eq).
	PrincipalId nullable.Type[string] `json:"principalId,omitempty"`

	// OData ID for `Principal` to bind to this entity
	Principal_ODataBind *string `json:"principal@odata.bind,omitempty"`

	// Fields inherited from PrivilegedAccessScheduleInstance

	// When the schedule instance ends, and is required.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// When this instance starts, and is required.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s PrivilegedAccessGroupEligibilityScheduleInstance) PrivilegedAccessScheduleInstance() BasePrivilegedAccessScheduleInstanceImpl {
	return BasePrivilegedAccessScheduleInstanceImpl{
		EndDateTime:   s.EndDateTime,
		StartDateTime: s.StartDateTime,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s PrivilegedAccessGroupEligibilityScheduleInstance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedAccessGroupEligibilityScheduleInstance{}

func (s PrivilegedAccessGroupEligibilityScheduleInstance) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedAccessGroupEligibilityScheduleInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedAccessGroupEligibilityScheduleInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedAccessGroupEligibilityScheduleInstance: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.privilegedAccessGroupEligibilityScheduleInstance"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedAccessGroupEligibilityScheduleInstance: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrivilegedAccessGroupEligibilityScheduleInstance{}

func (s *PrivilegedAccessGroupEligibilityScheduleInstance) UnmarshalJSON(bytes []byte) error {
	type alias PrivilegedAccessGroupEligibilityScheduleInstance
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into PrivilegedAccessGroupEligibilityScheduleInstance: %+v", err)
	}

	s.AccessId = decoded.AccessId
	s.EligibilityScheduleId = decoded.EligibilityScheduleId
	s.EndDateTime = decoded.EndDateTime
	s.Group = decoded.Group
	s.GroupId = decoded.GroupId
	s.Id = decoded.Id
	s.MemberType = decoded.MemberType
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PrincipalId = decoded.PrincipalId
	s.Principal_ODataBind = decoded.Principal_ODataBind
	s.StartDateTime = decoded.StartDateTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PrivilegedAccessGroupEligibilityScheduleInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'PrivilegedAccessGroupEligibilityScheduleInstance': %+v", err)
		}
		s.Principal = &impl
	}
	return nil
}
