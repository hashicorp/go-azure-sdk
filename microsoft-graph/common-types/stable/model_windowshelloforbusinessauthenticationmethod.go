package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethod = WindowsHelloForBusinessAuthenticationMethod{}

type WindowsHelloForBusinessAuthenticationMethod struct {
	// The date and time that this Windows Hello for Business key was registered.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The registered device on which this Windows Hello for Business key resides. Supports $expand. When you get a user's
	// Windows Hello for Business registration information, this property is returned only on a single GET and when you
	// specify ?$expand. For example, GET
	// /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
	Device *Device `json:"device,omitempty"`

	// The name of the device on which Windows Hello for Business is registered
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Key strength of this Windows Hello for Business key. Possible values are: normal, weak, unknown.
	KeyStrength *AuthenticationMethodKeyStrength `json:"keyStrength,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsHelloForBusinessAuthenticationMethod) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return BaseAuthenticationMethodImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WindowsHelloForBusinessAuthenticationMethod) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsHelloForBusinessAuthenticationMethod{}

func (s WindowsHelloForBusinessAuthenticationMethod) MarshalJSON() ([]byte, error) {
	type wrapper WindowsHelloForBusinessAuthenticationMethod
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsHelloForBusinessAuthenticationMethod: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsHelloForBusinessAuthenticationMethod: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsHelloForBusinessAuthenticationMethod"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsHelloForBusinessAuthenticationMethod: %+v", err)
	}

	return encoded, nil
}
