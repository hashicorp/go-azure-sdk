package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileContainedApp = MicrosoftStoreForBusinessContainedApp{}

type MicrosoftStoreForBusinessContainedApp struct {
	// The app user model ID of the contained app of a MicrosoftStoreForBusinessApp.
	AppUserModelId nullable.Type[string] `json:"appUserModelId,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s MicrosoftStoreForBusinessContainedApp) MobileContainedApp() BaseMobileContainedAppImpl {
	return BaseMobileContainedAppImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s MicrosoftStoreForBusinessContainedApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftStoreForBusinessContainedApp{}

func (s MicrosoftStoreForBusinessContainedApp) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftStoreForBusinessContainedApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftStoreForBusinessContainedApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftStoreForBusinessContainedApp: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.microsoftStoreForBusinessContainedApp"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftStoreForBusinessContainedApp: %+v", err)
	}

	return encoded, nil
}
