package functions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FunctionRetrieveDefaultDefinitionParameters = AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters{}

type AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters struct {
	BindingRetrievalProperties *AzureMachineLearningServiceFunctionBindingRetrievalProperties `json:"bindingRetrievalProperties,omitempty"`

	// Fields inherited from FunctionRetrieveDefaultDefinitionParameters
}

var _ json.Marshaler = AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters{}

func (s AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters) MarshalJSON() ([]byte, error) {
	type wrapper AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}
	decoded["bindingType"] = "Microsoft.MachineLearningServices"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters: %+v", err)
	}

	return encoded, nil
}
