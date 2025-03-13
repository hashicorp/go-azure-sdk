package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobOutput = UriFolderJobOutput{}

type UriFolderJobOutput struct {
	AssetName *string             `json:"assetName,omitempty"`
	Mode      *OutputDeliveryMode `json:"mode,omitempty"`
	Uri       *string             `json:"uri,omitempty"`

	// Fields inherited from JobOutput

	Description   *string       `json:"description,omitempty"`
	JobOutputType JobOutputType `json:"jobOutputType"`
}

func (s UriFolderJobOutput) JobOutput() BaseJobOutputImpl {
	return BaseJobOutputImpl{
		Description:   s.Description,
		JobOutputType: s.JobOutputType,
	}
}

var _ json.Marshaler = UriFolderJobOutput{}

func (s UriFolderJobOutput) MarshalJSON() ([]byte, error) {
	type wrapper UriFolderJobOutput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UriFolderJobOutput: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UriFolderJobOutput: %+v", err)
	}

	decoded["jobOutputType"] = "uri_folder"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UriFolderJobOutput: %+v", err)
	}

	return encoded, nil
}
