package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationEventListener = OnTokenIssuanceStartListener{}

type OnTokenIssuanceStartListener struct {
	// The handler to invoke when conditions are met for this onTokenIssuanceStartListener.
	Handler OnTokenIssuanceStartHandler `json:"handler"`

	// Fields inherited from AuthenticationEventListener

	// Indicates the authenticationEventListener is associated with an authenticationEventsFlow. Read-only.
	AuthenticationEventsFlowId nullable.Type[string] `json:"authenticationEventsFlowId,omitempty"`

	// The conditions on which this authenticationEventListener should trigger.
	Conditions *AuthenticationConditions `json:"conditions,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s OnTokenIssuanceStartListener) AuthenticationEventListener() BaseAuthenticationEventListenerImpl {
	return BaseAuthenticationEventListenerImpl{
		AuthenticationEventsFlowId: s.AuthenticationEventsFlowId,
		Conditions:                 s.Conditions,
		Id:                         s.Id,
		ODataId:                    s.ODataId,
		ODataType:                  s.ODataType,
	}
}

func (s OnTokenIssuanceStartListener) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnTokenIssuanceStartListener{}

func (s OnTokenIssuanceStartListener) MarshalJSON() ([]byte, error) {
	type wrapper OnTokenIssuanceStartListener
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnTokenIssuanceStartListener: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnTokenIssuanceStartListener: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.onTokenIssuanceStartListener"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnTokenIssuanceStartListener: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OnTokenIssuanceStartListener{}

func (s *OnTokenIssuanceStartListener) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		AuthenticationEventsFlowId nullable.Type[string]     `json:"authenticationEventsFlowId,omitempty"`
		Conditions                 *AuthenticationConditions `json:"conditions,omitempty"`
		Id                         *string                   `json:"id,omitempty"`
		ODataId                    *string                   `json:"@odata.id,omitempty"`
		ODataType                  *string                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AuthenticationEventsFlowId = decoded.AuthenticationEventsFlowId
	s.Conditions = decoded.Conditions
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OnTokenIssuanceStartListener into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["handler"]; ok {
		impl, err := UnmarshalOnTokenIssuanceStartHandlerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Handler' for 'OnTokenIssuanceStartListener': %+v", err)
		}
		s.Handler = impl
	}

	return nil
}
