package crossregionrestore

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossRegionRestoreRequest struct {
	CrossRegionRestoreAccessDetails CrrAccessToken `json:"crossRegionRestoreAccessDetails"`
	RestoreRequest                  RestoreRequest `json:"restoreRequest"`
}

var _ json.Unmarshaler = &CrossRegionRestoreRequest{}

func (s *CrossRegionRestoreRequest) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CrossRegionRestoreRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["crossRegionRestoreAccessDetails"]; ok {
		impl, err := unmarshalCrrAccessTokenImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CrossRegionRestoreAccessDetails' for 'CrossRegionRestoreRequest': %+v", err)
		}
		s.CrossRegionRestoreAccessDetails = impl
	}

	if v, ok := temp["restoreRequest"]; ok {
		impl, err := unmarshalRestoreRequestImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RestoreRequest' for 'CrossRegionRestoreRequest': %+v", err)
		}
		s.RestoreRequest = impl
	}
	return nil
}
