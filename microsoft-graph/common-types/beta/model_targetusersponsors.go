package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UserSet = TargetUserSponsors{}

type TargetUserSponsors struct {

	// Fields inherited from UserSet

	// For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
	IsBackup nullable.Type[bool] `json:"isBackup,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s TargetUserSponsors) UserSet() BaseUserSetImpl {
	return BaseUserSetImpl{
		IsBackup:  s.IsBackup,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TargetUserSponsors{}

func (s TargetUserSponsors) MarshalJSON() ([]byte, error) {
	type wrapper TargetUserSponsors
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TargetUserSponsors: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TargetUserSponsors: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.targetUserSponsors"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TargetUserSponsors: %+v", err)
	}

	return encoded, nil
}
