package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PackageInputPathBase = PackageInputPathVersion{}

type PackageInputPathVersion struct {
	ResourceName    *string `json:"resourceName,omitempty"`
	ResourceVersion *string `json:"resourceVersion,omitempty"`

	// Fields inherited from PackageInputPathBase
}

var _ json.Marshaler = PackageInputPathVersion{}

func (s PackageInputPathVersion) MarshalJSON() ([]byte, error) {
	type wrapper PackageInputPathVersion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PackageInputPathVersion: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PackageInputPathVersion: %+v", err)
	}
	decoded["inputPathType"] = "PathVersion"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PackageInputPathVersion: %+v", err)
	}

	return encoded, nil
}
