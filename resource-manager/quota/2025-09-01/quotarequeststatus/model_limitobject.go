package quotarequeststatus

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LimitJsonObject = LimitObject{}

type LimitObject struct {
	LimitType *QuotaLimitTypes `json:"limitType,omitempty"`
	Value     int64            `json:"value"`

	// Fields inherited from LimitJsonObject

	LimitObjectType LimitType `json:"limitObjectType"`
}

func (s LimitObject) LimitJsonObject() BaseLimitJsonObjectImpl {
	return BaseLimitJsonObjectImpl{
		LimitObjectType: s.LimitObjectType,
	}
}

var _ json.Marshaler = LimitObject{}

func (s LimitObject) MarshalJSON() ([]byte, error) {
	type wrapper LimitObject
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LimitObject: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LimitObject: %+v", err)
	}

	decoded["limitObjectType"] = "LimitValue"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LimitObject: %+v", err)
	}

	return encoded, nil
}
