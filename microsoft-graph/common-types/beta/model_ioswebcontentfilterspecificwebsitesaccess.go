package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosWebContentFilterBase = IosWebContentFilterSpecificWebsitesAccess{}

type IosWebContentFilterSpecificWebsitesAccess struct {
	// URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through
	// bookmarks. This collection can contain a maximum of 500 elements.
	SpecificWebsitesOnly *[]IosBookmark `json:"specificWebsitesOnly,omitempty"`

	// URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through
	// bookmarks. This collection can contain a maximum of 500 elements.
	WebsiteList *[]IosBookmark `json:"websiteList,omitempty"`

	// Fields inherited from IosWebContentFilterBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s IosWebContentFilterSpecificWebsitesAccess) IosWebContentFilterBase() BaseIosWebContentFilterBaseImpl {
	return BaseIosWebContentFilterBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosWebContentFilterSpecificWebsitesAccess{}

func (s IosWebContentFilterSpecificWebsitesAccess) MarshalJSON() ([]byte, error) {
	type wrapper IosWebContentFilterSpecificWebsitesAccess
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosWebContentFilterSpecificWebsitesAccess: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosWebContentFilterSpecificWebsitesAccess: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosWebContentFilterSpecificWebsitesAccess: %+v", err)
	}

	return encoded, nil
}
