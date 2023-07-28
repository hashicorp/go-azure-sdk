package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LabelingJobMediaProperties = LabelingJobTextProperties{}

type LabelingJobTextProperties struct {
	AnnotationType *TextAnnotationType `json:"annotationType,omitempty"`

	// Fields inherited from LabelingJobMediaProperties
}

var _ json.Marshaler = LabelingJobTextProperties{}

func (s LabelingJobTextProperties) MarshalJSON() ([]byte, error) {
	type wrapper LabelingJobTextProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LabelingJobTextProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LabelingJobTextProperties: %+v", err)
	}
	decoded["mediaType"] = "Text"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LabelingJobTextProperties: %+v", err)
	}

	return encoded, nil
}
