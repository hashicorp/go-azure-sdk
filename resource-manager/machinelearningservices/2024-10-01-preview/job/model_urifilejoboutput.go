package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobOutput = UriFileJobOutput{}

type UriFileJobOutput struct {
	AssetName *string             `json:"assetName,omitempty"`
	Mode      *OutputDeliveryMode `json:"mode,omitempty"`
	Uri       *string             `json:"uri,omitempty"`

	// Fields inherited from JobOutput

	Description   *string       `json:"description,omitempty"`
	JobOutputType JobOutputType `json:"jobOutputType"`
}

func (s UriFileJobOutput) JobOutput() BaseJobOutputImpl {
	return BaseJobOutputImpl{
		Description:   s.Description,
		JobOutputType: s.JobOutputType,
	}
}

var _ json.Marshaler = UriFileJobOutput{}

func (s UriFileJobOutput) MarshalJSON() ([]byte, error) {
	type wrapper UriFileJobOutput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UriFileJobOutput: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UriFileJobOutput: %+v", err)
	}

	decoded["jobOutputType"] = "uri_file"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UriFileJobOutput: %+v", err)
	}

	return encoded, nil
}
