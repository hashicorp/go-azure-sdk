package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ThreatAssessmentRequest = UrlAssessmentRequest{}

type UrlAssessmentRequest struct {
	// The URL string.
	Url *string `json:"url,omitempty"`

	// Fields inherited from ThreatAssessmentRequest

	Category *ThreatCategory `json:"category,omitempty"`

	// The content type of threat assessment. Possible values are: mail, url, file.
	ContentType *ThreatAssessmentContentType `json:"contentType,omitempty"`

	// The threat assessment request creator.
	CreatedBy IdentitySet `json:"createdBy"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	ExpectedAssessment *ThreatExpectedAssessment `json:"expectedAssessment,omitempty"`

	// The source of the threat assessment request. Possible values are: administrator.
	RequestSource *ThreatAssessmentRequestSource `json:"requestSource,omitempty"`

	// A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not
	// return this property unless you apply $expand on it.
	Results *[]ThreatAssessmentResult `json:"results,omitempty"`

	// The assessment process status. Possible values are: pending, completed.
	Status *ThreatAssessmentStatus `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s UrlAssessmentRequest) ThreatAssessmentRequest() BaseThreatAssessmentRequestImpl {
	return BaseThreatAssessmentRequestImpl{
		Category:           s.Category,
		ContentType:        s.ContentType,
		CreatedBy:          s.CreatedBy,
		CreatedDateTime:    s.CreatedDateTime,
		ExpectedAssessment: s.ExpectedAssessment,
		RequestSource:      s.RequestSource,
		Results:            s.Results,
		Status:             s.Status,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s UrlAssessmentRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UrlAssessmentRequest{}

func (s UrlAssessmentRequest) MarshalJSON() ([]byte, error) {
	type wrapper UrlAssessmentRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlAssessmentRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlAssessmentRequest: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.urlAssessmentRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlAssessmentRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UrlAssessmentRequest{}

func (s *UrlAssessmentRequest) UnmarshalJSON(bytes []byte) error {
	type alias UrlAssessmentRequest
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into UrlAssessmentRequest: %+v", err)
	}

	s.Category = decoded.Category
	s.ContentType = decoded.ContentType
	s.CreatedDateTime = decoded.CreatedDateTime
	s.ExpectedAssessment = decoded.ExpectedAssessment
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RequestSource = decoded.RequestSource
	s.Results = decoded.Results
	s.Status = decoded.Status
	s.Url = decoded.Url

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UrlAssessmentRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'UrlAssessmentRequest': %+v", err)
		}
		s.CreatedBy = impl
	}
	return nil
}
