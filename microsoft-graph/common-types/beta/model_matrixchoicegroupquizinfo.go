package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ QuizInfo = MatrixChoiceGroupQuizInfo{}

type MatrixChoiceGroupQuizInfo struct {

	// Fields inherited from QuizInfo

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MatrixChoiceGroupQuizInfo) QuizInfo() BaseQuizInfoImpl {
	return BaseQuizInfoImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MatrixChoiceGroupQuizInfo{}

func (s MatrixChoiceGroupQuizInfo) MarshalJSON() ([]byte, error) {
	type wrapper MatrixChoiceGroupQuizInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MatrixChoiceGroupQuizInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MatrixChoiceGroupQuizInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.matrixChoiceGroupQuizInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MatrixChoiceGroupQuizInfo: %+v", err)
	}

	return encoded, nil
}
