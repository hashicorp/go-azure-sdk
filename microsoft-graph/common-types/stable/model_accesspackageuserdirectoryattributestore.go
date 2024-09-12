package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessPackageResourceAttributeDestination = AccessPackageUserDirectoryAttributeStore{}

type AccessPackageUserDirectoryAttributeStore struct {

	// Fields inherited from AccessPackageResourceAttributeDestination

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AccessPackageUserDirectoryAttributeStore) AccessPackageResourceAttributeDestination() BaseAccessPackageResourceAttributeDestinationImpl {
	return BaseAccessPackageResourceAttributeDestinationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageUserDirectoryAttributeStore{}

func (s AccessPackageUserDirectoryAttributeStore) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageUserDirectoryAttributeStore
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageUserDirectoryAttributeStore: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageUserDirectoryAttributeStore: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.accessPackageUserDirectoryAttributeStore"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageUserDirectoryAttributeStore: %+v", err)
	}

	return encoded, nil
}
