package discoveryrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DiscoveryRuleSpecification = ResourceGraphQuerySpecification{}

type ResourceGraphQuerySpecification struct {
	ResourceGraphQuery string `json:"resourceGraphQuery"`

	// Fields inherited from DiscoveryRuleSpecification

	Kind DiscoveryRuleKind `json:"kind"`
}

func (s ResourceGraphQuerySpecification) DiscoveryRuleSpecification() BaseDiscoveryRuleSpecificationImpl {
	return BaseDiscoveryRuleSpecificationImpl{
		Kind: s.Kind,
	}
}

var _ json.Marshaler = ResourceGraphQuerySpecification{}

func (s ResourceGraphQuerySpecification) MarshalJSON() ([]byte, error) {
	type wrapper ResourceGraphQuerySpecification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResourceGraphQuerySpecification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceGraphQuerySpecification: %+v", err)
	}

	decoded["kind"] = "ResourceGraphQuery"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResourceGraphQuerySpecification: %+v", err)
	}

	return encoded, nil
}
