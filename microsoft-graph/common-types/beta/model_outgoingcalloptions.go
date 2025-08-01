package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallOptions = OutgoingCallOptions{}

type OutgoingCallOptions struct {

	// Fields inherited from CallOptions

	// Indicates whether to hide the app after the call is escalated.
	HideBotAfterEscalation nullable.Type[bool] `json:"hideBotAfterEscalation,omitempty"`

	// Indicates whether content sharing notifications should be enabled for the call.
	IsContentSharingNotificationEnabled nullable.Type[bool] `json:"isContentSharingNotificationEnabled,omitempty"`

	// Indicates whether delta roster is enabled for the call.
	IsDeltaRosterEnabled nullable.Type[bool] `json:"isDeltaRosterEnabled,omitempty"`

	// Indicates whether delta roster filtering by participant interactivity is enabled.
	IsInteractiveRosterEnabled nullable.Type[bool] `json:"isInteractiveRosterEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OutgoingCallOptions) CallOptions() BaseCallOptionsImpl {
	return BaseCallOptionsImpl{
		HideBotAfterEscalation:              s.HideBotAfterEscalation,
		IsContentSharingNotificationEnabled: s.IsContentSharingNotificationEnabled,
		IsDeltaRosterEnabled:                s.IsDeltaRosterEnabled,
		IsInteractiveRosterEnabled:          s.IsInteractiveRosterEnabled,
		ODataId:                             s.ODataId,
		ODataType:                           s.ODataType,
	}
}

var _ json.Marshaler = OutgoingCallOptions{}

func (s OutgoingCallOptions) MarshalJSON() ([]byte, error) {
	type wrapper OutgoingCallOptions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OutgoingCallOptions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OutgoingCallOptions: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.outgoingCallOptions"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OutgoingCallOptions: %+v", err)
	}

	return encoded, nil
}
