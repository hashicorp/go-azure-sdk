package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UserExperienceAnalyticsInsightValue = InsightValueDouble{}

type InsightValueDouble struct {

	// Fields inherited from UserExperienceAnalyticsInsightValue

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s InsightValueDouble) UserExperienceAnalyticsInsightValue() BaseUserExperienceAnalyticsInsightValueImpl {
	return BaseUserExperienceAnalyticsInsightValueImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InsightValueDouble{}

func (s InsightValueDouble) MarshalJSON() ([]byte, error) {
	type wrapper InsightValueDouble
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InsightValueDouble: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InsightValueDouble: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.insightValueDouble"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InsightValueDouble: %+v", err)
	}

	return encoded, nil
}
