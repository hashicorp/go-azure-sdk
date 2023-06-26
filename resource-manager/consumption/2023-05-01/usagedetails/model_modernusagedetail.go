package usagedetails

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UsageDetail = ModernUsageDetail{}

type ModernUsageDetail struct {
	Properties ModernUsageDetailProperties `json:"properties"`

	// Fields inherited from UsageDetail
	Etag *string            `json:"etag,omitempty"`
	Id   *string            `json:"id,omitempty"`
	Name *string            `json:"name,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	Type *string            `json:"type,omitempty"`
}

var _ json.Marshaler = ModernUsageDetail{}

func (s ModernUsageDetail) MarshalJSON() ([]byte, error) {
	type wrapper ModernUsageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ModernUsageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ModernUsageDetail: %+v", err)
	}
	decoded["kind"] = "modern"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ModernUsageDetail: %+v", err)
	}

	return encoded, nil
}
