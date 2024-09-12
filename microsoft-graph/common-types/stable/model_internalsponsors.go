package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectSet = InternalSponsors{}

type InternalSponsors struct {

	// Fields inherited from SubjectSet

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s InternalSponsors) SubjectSet() BaseSubjectSetImpl {
	return BaseSubjectSetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InternalSponsors{}

func (s InternalSponsors) MarshalJSON() ([]byte, error) {
	type wrapper InternalSponsors
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InternalSponsors: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InternalSponsors: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.internalSponsors"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InternalSponsors: %+v", err)
	}

	return encoded, nil
}
