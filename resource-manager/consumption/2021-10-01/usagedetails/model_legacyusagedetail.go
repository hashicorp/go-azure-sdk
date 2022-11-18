package usagedetails

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UsageDetail = LegacyUsageDetail{}

type LegacyUsageDetail struct {
	Properties LegacyUsageDetailProperties `json:"properties"`

	// Fields inherited from UsageDetail
	Etag *string            `json:"etag,omitempty"`
	Id   *string            `json:"id,omitempty"`
	Name *string            `json:"name,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	Type *string            `json:"type,omitempty"`
}

var _ json.Marshaler = LegacyUsageDetail{}

func (s LegacyUsageDetail) MarshalJSON() ([]byte, error) {
	type wrapper LegacyUsageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LegacyUsageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LegacyUsageDetail: %+v", err)
	}
	decoded["kind"] = "legacy"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LegacyUsageDetail: %+v", err)
	}

	return encoded, nil
}
