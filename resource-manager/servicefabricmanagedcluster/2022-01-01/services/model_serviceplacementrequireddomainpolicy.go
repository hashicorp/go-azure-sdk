package services

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServicePlacementPolicy = ServicePlacementRequiredDomainPolicy{}

type ServicePlacementRequiredDomainPolicy struct {
	DomainName string `json:"domainName"`

	// Fields inherited from ServicePlacementPolicy

	Type ServicePlacementPolicyType `json:"type"`
}

func (s ServicePlacementRequiredDomainPolicy) ServicePlacementPolicy() BaseServicePlacementPolicyImpl {
	return BaseServicePlacementPolicyImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ServicePlacementRequiredDomainPolicy{}

func (s ServicePlacementRequiredDomainPolicy) MarshalJSON() ([]byte, error) {
	type wrapper ServicePlacementRequiredDomainPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServicePlacementRequiredDomainPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServicePlacementRequiredDomainPolicy: %+v", err)
	}

	decoded["type"] = "RequiredDomain"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServicePlacementRequiredDomainPolicy: %+v", err)
	}

	return encoded, nil
}
