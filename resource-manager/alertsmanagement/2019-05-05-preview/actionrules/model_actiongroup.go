package actionrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ActionRuleProperties = ActionGroup{}

type ActionGroup struct {
	ActionGroupId *string `json:"actionGroupId,omitempty"`

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

var _ json.Marshaler = ActionGroup{}

func (s ActionGroup) MarshalJSON() ([]byte, error) {
	type wrapper ActionGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ActionGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ActionGroup: %+v", err)
	}
	decoded["type"] = "ActionGroup"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ActionGroup: %+v", err)
	}

	return encoded, nil
}
