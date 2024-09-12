package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = EncryptedGcpStorageBucketFinding{}

type EncryptedGcpStorageBucketFinding struct {
	Accessibility       *GcpAccessType               `json:"accessibility,omitempty"`
	EncryptionManagedBy *GcpEncryption               `json:"encryptionManagedBy,omitempty"`
	StorageBucket       *AuthorizationSystemResource `json:"storageBucket,omitempty"`

	// Fields inherited from Finding

	// Defines when the finding was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s EncryptedGcpStorageBucketFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s EncryptedGcpStorageBucketFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EncryptedGcpStorageBucketFinding{}

func (s EncryptedGcpStorageBucketFinding) MarshalJSON() ([]byte, error) {
	type wrapper EncryptedGcpStorageBucketFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EncryptedGcpStorageBucketFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EncryptedGcpStorageBucketFinding: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.encryptedGcpStorageBucketFinding"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EncryptedGcpStorageBucketFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EncryptedGcpStorageBucketFinding{}

func (s *EncryptedGcpStorageBucketFinding) UnmarshalJSON(bytes []byte) error {
	type alias EncryptedGcpStorageBucketFinding
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into EncryptedGcpStorageBucketFinding: %+v", err)
	}

	s.Accessibility = decoded.Accessibility
	s.CreatedDateTime = decoded.CreatedDateTime
	s.EncryptionManagedBy = decoded.EncryptionManagedBy
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EncryptedGcpStorageBucketFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["storageBucket"]; ok {
		impl, err := UnmarshalAuthorizationSystemResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StorageBucket' for 'EncryptedGcpStorageBucketFinding': %+v", err)
		}
		s.StorageBucket = &impl
	}
	return nil
}
