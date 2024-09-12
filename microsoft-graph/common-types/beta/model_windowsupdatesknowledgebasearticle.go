package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesKnowledgeBaseArticle{}

type WindowsUpdatesKnowledgeBaseArticle struct {
	// The URL of the knowledge base article. Read-only.
	Url *string `json:"url,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsUpdatesKnowledgeBaseArticle) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesKnowledgeBaseArticle{}

func (s WindowsUpdatesKnowledgeBaseArticle) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesKnowledgeBaseArticle
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesKnowledgeBaseArticle: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesKnowledgeBaseArticle: %+v", err)
	}

	delete(decoded, "url")
	decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.knowledgeBaseArticle"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesKnowledgeBaseArticle: %+v", err)
	}

	return encoded, nil
}
