package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ThumbnailSet{}

type ThumbnailSet struct {
	// A 1920x1920 scaled thumbnail.
	Large *Thumbnail `json:"large,omitempty"`

	// A 176x176 scaled thumbnail.
	Medium *Thumbnail `json:"medium,omitempty"`

	// A 48x48 cropped thumbnail.
	Small *Thumbnail `json:"small,omitempty"`

	// A custom thumbnail image or the original image used to generate other thumbnails.
	Source *Thumbnail `json:"source,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ThumbnailSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ThumbnailSet{}

func (s ThumbnailSet) MarshalJSON() ([]byte, error) {
	type wrapper ThumbnailSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ThumbnailSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ThumbnailSet: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.thumbnailSet"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ThumbnailSet: %+v", err)
	}

	return encoded, nil
}
