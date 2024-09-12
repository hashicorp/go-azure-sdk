package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdateCatalogItem = WindowsFeatureUpdateCatalogItem{}

type WindowsFeatureUpdateCatalogItem struct {
	// The feature update version
	Version *string `json:"version,omitempty"`

	// Fields inherited from WindowsUpdateCatalogItem

	// The display name for the catalog item.
	DisplayName *string `json:"displayName,omitempty"`

	// The last supported date for a catalog item
	EndOfSupportDate nullable.Type[string] `json:"endOfSupportDate,omitempty"`

	// The date the catalog item was released
	ReleaseDateTime *string `json:"releaseDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsFeatureUpdateCatalogItem) WindowsUpdateCatalogItem() BaseWindowsUpdateCatalogItemImpl {
	return BaseWindowsUpdateCatalogItemImpl{
		DisplayName:      s.DisplayName,
		EndOfSupportDate: s.EndOfSupportDate,
		ReleaseDateTime:  s.ReleaseDateTime,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s WindowsFeatureUpdateCatalogItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsFeatureUpdateCatalogItem{}

func (s WindowsFeatureUpdateCatalogItem) MarshalJSON() ([]byte, error) {
	type wrapper WindowsFeatureUpdateCatalogItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsFeatureUpdateCatalogItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsFeatureUpdateCatalogItem: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsFeatureUpdateCatalogItem"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsFeatureUpdateCatalogItem: %+v", err)
	}

	return encoded, nil
}
