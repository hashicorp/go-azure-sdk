package labelingjob

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LabelingJobMediaProperties = LabelingJobImageProperties{}

type LabelingJobImageProperties struct {
	AnnotationType *ImageAnnotationType `json:"annotationType,omitempty"`

	// Fields inherited from LabelingJobMediaProperties
}

var _ json.Marshaler = LabelingJobImageProperties{}

func (s LabelingJobImageProperties) MarshalJSON() ([]byte, error) {
	type wrapper LabelingJobImageProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LabelingJobImageProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LabelingJobImageProperties: %+v", err)
	}
	decoded["mediaType"] = "Image"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LabelingJobImageProperties: %+v", err)
	}

	return encoded, nil
}
