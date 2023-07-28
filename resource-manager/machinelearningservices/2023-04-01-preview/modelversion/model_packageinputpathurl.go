package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PackageInputPathBase = PackageInputPathUrl{}

type PackageInputPathUrl struct {
	Url *string `json:"url,omitempty"`

	// Fields inherited from PackageInputPathBase
}

var _ json.Marshaler = PackageInputPathUrl{}

func (s PackageInputPathUrl) MarshalJSON() ([]byte, error) {
	type wrapper PackageInputPathUrl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PackageInputPathUrl: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PackageInputPathUrl: %+v", err)
	}
	decoded["inputPathType"] = "Url"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PackageInputPathUrl: %+v", err)
	}

	return encoded, nil
}
