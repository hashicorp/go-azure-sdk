package functions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FunctionRetrieveDefaultDefinitionParameters = CSharpFunctionRetrieveDefaultDefinitionParameters{}

type CSharpFunctionRetrieveDefaultDefinitionParameters struct {
	BindingRetrievalProperties *CSharpFunctionBindingRetrievalProperties `json:"bindingRetrievalProperties,omitempty"`

	// Fields inherited from FunctionRetrieveDefaultDefinitionParameters

	BindingType string `json:"bindingType"`
}

func (s CSharpFunctionRetrieveDefaultDefinitionParameters) FunctionRetrieveDefaultDefinitionParameters() BaseFunctionRetrieveDefaultDefinitionParametersImpl {
	return BaseFunctionRetrieveDefaultDefinitionParametersImpl{
		BindingType: s.BindingType,
	}
}

var _ json.Marshaler = CSharpFunctionRetrieveDefaultDefinitionParameters{}

func (s CSharpFunctionRetrieveDefaultDefinitionParameters) MarshalJSON() ([]byte, error) {
	type wrapper CSharpFunctionRetrieveDefaultDefinitionParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CSharpFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CSharpFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}

	decoded["bindingType"] = "Microsoft.StreamAnalytics/CLRUdf"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CSharpFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}

	return encoded, nil
}
