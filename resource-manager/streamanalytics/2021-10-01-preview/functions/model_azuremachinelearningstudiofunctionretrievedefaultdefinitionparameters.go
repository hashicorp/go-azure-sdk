package functions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FunctionRetrieveDefaultDefinitionParameters = AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters{}

type AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters struct {
	BindingRetrievalProperties *AzureMachineLearningStudioFunctionBindingRetrievalProperties `json:"bindingRetrievalProperties,omitempty"`

	// Fields inherited from FunctionRetrieveDefaultDefinitionParameters

	BindingType string `json:"bindingType"`
}

func (s AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters) FunctionRetrieveDefaultDefinitionParameters() BaseFunctionRetrieveDefaultDefinitionParametersImpl {
	return BaseFunctionRetrieveDefaultDefinitionParametersImpl{
		BindingType: s.BindingType,
	}
}

var _ json.Marshaler = AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters{}

func (s AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters) MarshalJSON() ([]byte, error) {
	type wrapper AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}

	decoded["bindingType"] = "Microsoft.MachineLearning/WebService"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}

	return encoded, nil
}
