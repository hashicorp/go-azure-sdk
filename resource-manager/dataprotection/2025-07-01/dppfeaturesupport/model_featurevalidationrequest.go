package dppfeaturesupport

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FeatureValidationRequestBase = FeatureValidationRequest{}

type FeatureValidationRequest struct {
	FeatureName *string      `json:"featureName,omitempty"`
	FeatureType *FeatureType `json:"featureType,omitempty"`

	// Fields inherited from FeatureValidationRequestBase

	ObjectType string `json:"objectType"`
}

func (s FeatureValidationRequest) FeatureValidationRequestBase() BaseFeatureValidationRequestBaseImpl {
	return BaseFeatureValidationRequestBaseImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = FeatureValidationRequest{}

func (s FeatureValidationRequest) MarshalJSON() ([]byte, error) {
	type wrapper FeatureValidationRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FeatureValidationRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureValidationRequest: %+v", err)
	}

	decoded["objectType"] = "FeatureValidationRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FeatureValidationRequest: %+v", err)
	}

	return encoded, nil
}
