package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionUnitsBulkJobBase interface {
	Entity
	ProtectionUnitsBulkJobBase() BaseProtectionUnitsBulkJobBaseImpl
}

var _ ProtectionUnitsBulkJobBase = BaseProtectionUnitsBulkJobBaseImpl{}

type BaseProtectionUnitsBulkJobBaseImpl struct {
	// The identity of person who created the job.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the job.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The name of the protection units bulk addition job.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Error details containing resource resolution failures, if any.
	Error *PublicError `json:"error,omitempty"`

	// The identity of the person who last modified the job.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification made to the job.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	Status *ProtectionUnitsBulkJobStatus `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseProtectionUnitsBulkJobBaseImpl) ProtectionUnitsBulkJobBase() BaseProtectionUnitsBulkJobBaseImpl {
	return s
}

func (s BaseProtectionUnitsBulkJobBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ProtectionUnitsBulkJobBase = RawProtectionUnitsBulkJobBaseImpl{}

// RawProtectionUnitsBulkJobBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProtectionUnitsBulkJobBaseImpl struct {
	protectionUnitsBulkJobBase BaseProtectionUnitsBulkJobBaseImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawProtectionUnitsBulkJobBaseImpl) ProtectionUnitsBulkJobBase() BaseProtectionUnitsBulkJobBaseImpl {
	return s.protectionUnitsBulkJobBase
}

func (s RawProtectionUnitsBulkJobBaseImpl) Entity() BaseEntityImpl {
	return s.protectionUnitsBulkJobBase.Entity()
}

var _ json.Marshaler = BaseProtectionUnitsBulkJobBaseImpl{}

func (s BaseProtectionUnitsBulkJobBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseProtectionUnitsBulkJobBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseProtectionUnitsBulkJobBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseProtectionUnitsBulkJobBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.protectionUnitsBulkJobBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseProtectionUnitsBulkJobBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseProtectionUnitsBulkJobBaseImpl{}

func (s *BaseProtectionUnitsBulkJobBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string]         `json:"createdDateTime,omitempty"`
		DisplayName          nullable.Type[string]         `json:"displayName,omitempty"`
		Error                *PublicError                  `json:"error,omitempty"`
		LastModifiedDateTime nullable.Type[string]         `json:"lastModifiedDateTime,omitempty"`
		Status               *ProtectionUnitsBulkJobStatus `json:"status,omitempty"`
		Id                   *string                       `json:"id,omitempty"`
		ODataId              *string                       `json:"@odata.id,omitempty"`
		ODataType            *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Error = decoded.Error
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseProtectionUnitsBulkJobBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseProtectionUnitsBulkJobBaseImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseProtectionUnitsBulkJobBaseImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalProtectionUnitsBulkJobBaseImplementation(input []byte) (ProtectionUnitsBulkJobBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectionUnitsBulkJobBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.driveProtectionUnitsBulkAdditionJob") {
		var out DriveProtectionUnitsBulkAdditionJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveProtectionUnitsBulkAdditionJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxProtectionUnitsBulkAdditionJob") {
		var out MailboxProtectionUnitsBulkAdditionJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxProtectionUnitsBulkAdditionJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteProtectionUnitsBulkAdditionJob") {
		var out SiteProtectionUnitsBulkAdditionJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteProtectionUnitsBulkAdditionJob: %+v", err)
		}
		return out, nil
	}

	var parent BaseProtectionUnitsBulkJobBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProtectionUnitsBulkJobBaseImpl: %+v", err)
	}

	return RawProtectionUnitsBulkJobBaseImpl{
		protectionUnitsBulkJobBase: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
