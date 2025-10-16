package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorTemplateAllocation = AkriConnectorTemplateBucketizedAllocation{}

type AkriConnectorTemplateBucketizedAllocation struct {
	BucketSize int64 `json:"bucketSize"`

	// Fields inherited from AkriConnectorTemplateAllocation

	Policy AkriConnectorTemplateAllocationPolicy `json:"policy"`
}

func (s AkriConnectorTemplateBucketizedAllocation) AkriConnectorTemplateAllocation() BaseAkriConnectorTemplateAllocationImpl {
	return BaseAkriConnectorTemplateAllocationImpl{
		Policy: s.Policy,
	}
}

var _ json.Marshaler = AkriConnectorTemplateBucketizedAllocation{}

func (s AkriConnectorTemplateBucketizedAllocation) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorTemplateBucketizedAllocation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorTemplateBucketizedAllocation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorTemplateBucketizedAllocation: %+v", err)
	}

	decoded["policy"] = "Bucketized"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorTemplateBucketizedAllocation: %+v", err)
	}

	return encoded, nil
}
