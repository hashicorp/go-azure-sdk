package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AttachmentSession{}

type AttachmentSession struct {
	// The content streams that are uploaded.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The date and time in UTC when the upload session will expire. The complete file must be uploaded before this
	// expiration time is reached.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// Indicates a single value {start} that represents the location in the file where the next upload should begin.
	NextExpectedRanges *[]string `json:"nextExpectedRanges,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AttachmentSession) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AttachmentSession{}

func (s AttachmentSession) MarshalJSON() ([]byte, error) {
	type wrapper AttachmentSession
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AttachmentSession: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AttachmentSession: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.attachmentSession"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AttachmentSession: %+v", err)
	}

	return encoded, nil
}
