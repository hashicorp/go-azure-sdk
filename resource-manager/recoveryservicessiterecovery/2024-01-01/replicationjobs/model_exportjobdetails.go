package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobDetails = ExportJobDetails{}

type ExportJobDetails struct {
	BlobUri  *string `json:"blobUri,omitempty"`
	SasToken *string `json:"sasToken,omitempty"`

	// Fields inherited from JobDetails
	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
}

var _ json.Marshaler = ExportJobDetails{}

func (s ExportJobDetails) MarshalJSON() ([]byte, error) {
	type wrapper ExportJobDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExportJobDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExportJobDetails: %+v", err)
	}
	decoded["instanceType"] = "ExportJobDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExportJobDetails: %+v", err)
	}

	return encoded, nil
}
