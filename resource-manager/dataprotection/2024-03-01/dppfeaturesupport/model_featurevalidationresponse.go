package dppfeaturesupport

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FeatureValidationResponseBase = FeatureValidationResponse{}

type FeatureValidationResponse struct {
	FeatureType *FeatureType        `json:"featureType,omitempty"`
	Features    *[]SupportedFeature `json:"features,omitempty"`

	// Fields inherited from FeatureValidationResponseBase

	ObjectType string `json:"objectType"`
}

func (s FeatureValidationResponse) FeatureValidationResponseBase() BaseFeatureValidationResponseBaseImpl {
	return BaseFeatureValidationResponseBaseImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = FeatureValidationResponse{}

func (s FeatureValidationResponse) MarshalJSON() ([]byte, error) {
	type wrapper FeatureValidationResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FeatureValidationResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureValidationResponse: %+v", err)
	}

	decoded["objectType"] = "FeatureValidationResponse"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FeatureValidationResponse: %+v", err)
	}

	return encoded, nil
}
