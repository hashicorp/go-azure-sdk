package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectRightsRequestMailboxLocation = SubjectRightsRequestAllMailboxLocation{}

type SubjectRightsRequestAllMailboxLocation struct {

	// Fields inherited from SubjectRightsRequestMailboxLocation

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SubjectRightsRequestAllMailboxLocation) SubjectRightsRequestMailboxLocation() BaseSubjectRightsRequestMailboxLocationImpl {
	return BaseSubjectRightsRequestMailboxLocationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SubjectRightsRequestAllMailboxLocation{}

func (s SubjectRightsRequestAllMailboxLocation) MarshalJSON() ([]byte, error) {
	type wrapper SubjectRightsRequestAllMailboxLocation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SubjectRightsRequestAllMailboxLocation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SubjectRightsRequestAllMailboxLocation: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.subjectRightsRequestAllMailboxLocation"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SubjectRightsRequestAllMailboxLocation: %+v", err)
	}

	return encoded, nil
}
