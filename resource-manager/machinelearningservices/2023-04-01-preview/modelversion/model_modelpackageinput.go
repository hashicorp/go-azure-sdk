package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelPackageInput struct {
	InputType PackageInputType          `json:"inputType"`
	Mode      *PackageInputDeliveryMode `json:"mode,omitempty"`
	MountPath *string                   `json:"mountPath,omitempty"`
	Path      PackageInputPathBase      `json:"path"`
}

var _ json.Unmarshaler = &ModelPackageInput{}

func (s *ModelPackageInput) UnmarshalJSON(bytes []byte) error {
	type alias ModelPackageInput
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ModelPackageInput: %+v", err)
	}

	s.InputType = decoded.InputType
	s.Mode = decoded.Mode
	s.MountPath = decoded.MountPath

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ModelPackageInput into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["path"]; ok {
		impl, err := unmarshalPackageInputPathBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Path' for 'ModelPackageInput': %+v", err)
		}
		s.Path = impl
	}
	return nil
}
