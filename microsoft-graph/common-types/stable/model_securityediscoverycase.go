package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityCase = SecurityEdiscoveryCase{}

type SecurityEdiscoveryCase struct {
	// The user who closed the case.
	ClosedBy IdentitySet `json:"closedBy"`

	// The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ClosedDateTime nullable.Type[string] `json:"closedDateTime,omitempty"`

	// Returns a list of case ediscoveryCustodian objects for this case.
	Custodians *[]SecurityEdiscoveryCustodian `json:"custodians,omitempty"`

	// The external case number for customer reference.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// Returns a list of case ediscoveryNoncustodialDataSource objects for this case.
	NoncustodialDataSources *[]SecurityEdiscoveryNoncustodialDataSource `json:"noncustodialDataSources,omitempty"`

	// Returns a list of case caseOperation objects for this case.
	Operations *[]SecurityCaseOperation `json:"operations,omitempty"`

	// Returns a list of eDiscoveryReviewSet objects in the case.
	ReviewSets *[]SecurityEdiscoveryReviewSet `json:"reviewSets,omitempty"`

	// Returns a list of eDiscoverySearch objects associated with this case.
	Searches *[]SecurityEdiscoverySearch `json:"searches,omitempty"`

	// Returns a list of eDIscoverySettings objects in the case.
	Settings *SecurityEdiscoveryCaseSettings `json:"settings,omitempty"`

	// Returns a list of ediscoveryReviewTag objects associated to this case.
	Tags *[]SecurityEdiscoveryReviewTag `json:"tags,omitempty"`

	// Fields inherited from SecurityCase

	CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
	Description          nullable.Type[string] `json:"description,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedBy       IdentitySet           `json:"lastModifiedBy"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
	Status               *SecurityCaseStatus   `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityEdiscoveryCase) SecurityCase() BaseSecurityCaseImpl {
	return BaseSecurityCaseImpl{
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityEdiscoveryCase) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryCase{}

func (s SecurityEdiscoveryCase) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryCase
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryCase: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryCase: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryCase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryCase: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryCase{}

func (s *SecurityEdiscoveryCase) UnmarshalJSON(bytes []byte) error {
	type alias SecurityEdiscoveryCase
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SecurityEdiscoveryCase: %+v", err)
	}

	s.ClosedDateTime = decoded.ClosedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Custodians = decoded.Custodians
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.NoncustodialDataSources = decoded.NoncustodialDataSources
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ReviewSets = decoded.ReviewSets
	s.Searches = decoded.Searches
	s.Settings = decoded.Settings
	s.Status = decoded.Status
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityEdiscoveryCase into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["closedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ClosedBy' for 'SecurityEdiscoveryCase': %+v", err)
		}
		s.ClosedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SecurityEdiscoveryCase': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityCaseOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityCaseOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'SecurityEdiscoveryCase': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}
	return nil
}
