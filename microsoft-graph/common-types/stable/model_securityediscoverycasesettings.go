package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityEdiscoveryCaseSettings{}

type SecurityEdiscoveryCaseSettings struct {
	// The OCR (Optical Character Recognition) settings for the case.
	Ocr *SecurityOcrSettings `json:"ocr,omitempty"`

	// The redundancy (near duplicate and email threading) detection settings for the case.
	RedundancyDetection *SecurityRedundancyDetectionSettings `json:"redundancyDetection,omitempty"`

	// The Topic Modeling (Themes) settings for the case.
	TopicModeling *SecurityTopicModelingSettings `json:"topicModeling,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s SecurityEdiscoveryCaseSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryCaseSettings{}

func (s SecurityEdiscoveryCaseSettings) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryCaseSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryCaseSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryCaseSettings: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryCaseSettings"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryCaseSettings: %+v", err)
	}

	return encoded, nil
}
