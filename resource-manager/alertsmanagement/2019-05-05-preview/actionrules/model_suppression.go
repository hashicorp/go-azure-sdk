package actionrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ActionRuleProperties = Suppression{}

type Suppression struct {
	SuppressionConfig *SuppressionConfig `json:"suppressionConfig,omitempty"`

	// Fields inherited from ActionRuleProperties
	Conditions     *Conditions       `json:"conditions,omitempty"`
	CreatedAt      *string           `json:"createdAt,omitempty"`
	CreatedBy      *string           `json:"createdBy,omitempty"`
	Description    *string           `json:"description,omitempty"`
	LastModifiedAt *string           `json:"lastModifiedAt,omitempty"`
	LastModifiedBy *string           `json:"lastModifiedBy,omitempty"`
	Scope          *Scope            `json:"scope,omitempty"`
	Status         *ActionRuleStatus `json:"status,omitempty"`
}

var _ json.Marshaler = Suppression{}

func (s Suppression) MarshalJSON() ([]byte, error) {
	type wrapper Suppression
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Suppression: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Suppression: %+v", err)
	}
	decoded["type"] = "Suppression"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Suppression: %+v", err)
	}

	return encoded, nil
}
