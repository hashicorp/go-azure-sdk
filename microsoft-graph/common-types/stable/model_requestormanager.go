package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectSet = RequestorManager{}

type RequestorManager struct {
	// The hierarchical level of the manager with respect to the requestor. For example, the direct manager of a requestor
	// would have a managerLevel of 1, while the manager of the requestor's manager would have a managerLevel of 2. Default
	// value for managerLevel is 1. Possible values for this property range from 1 to 2.
	ManagerLevel nullable.Type[int64] `json:"managerLevel,omitempty"`

	// Fields inherited from SubjectSet

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s RequestorManager) SubjectSet() BaseSubjectSetImpl {
	return BaseSubjectSetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RequestorManager{}

func (s RequestorManager) MarshalJSON() ([]byte, error) {
	type wrapper RequestorManager
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RequestorManager: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RequestorManager: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.requestorManager"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RequestorManager: %+v", err)
	}

	return encoded, nil
}
