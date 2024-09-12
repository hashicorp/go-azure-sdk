package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OmaSetting = OmaSettingFloatingPoint{}

type OmaSettingFloatingPoint struct {

	// Fields inherited from OmaSetting

	// Description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display Name.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates whether the value field is encrypted. This property is read-only.
	IsEncrypted *bool `json:"isEncrypted,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// OMA.
	OmaUri *string `json:"omaUri,omitempty"`

	// ReferenceId for looking up secret for decryption. This property is read-only.
	SecretReferenceValueId nullable.Type[string] `json:"secretReferenceValueId,omitempty"`
}

func (s OmaSettingFloatingPoint) OmaSetting() BaseOmaSettingImpl {
	return BaseOmaSettingImpl{
		Description:            s.Description,
		DisplayName:            s.DisplayName,
		IsEncrypted:            s.IsEncrypted,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
		OmaUri:                 s.OmaUri,
		SecretReferenceValueId: s.SecretReferenceValueId,
	}
}

var _ json.Marshaler = OmaSettingFloatingPoint{}

func (s OmaSettingFloatingPoint) MarshalJSON() ([]byte, error) {
	type wrapper OmaSettingFloatingPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OmaSettingFloatingPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OmaSettingFloatingPoint: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.omaSettingFloatingPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OmaSettingFloatingPoint: %+v", err)
	}

	return encoded, nil
}
