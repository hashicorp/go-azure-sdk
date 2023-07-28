package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PackageInputPathBase = PackageInputPathId{}

type PackageInputPathId struct {
	ResourceId *string `json:"resourceId,omitempty"`

	// Fields inherited from PackageInputPathBase
}

var _ json.Marshaler = PackageInputPathId{}

func (s PackageInputPathId) MarshalJSON() ([]byte, error) {
	type wrapper PackageInputPathId
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PackageInputPathId: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PackageInputPathId: %+v", err)
	}
	decoded["inputPathType"] = "PathId"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PackageInputPathId: %+v", err)
	}

	return encoded, nil
}
