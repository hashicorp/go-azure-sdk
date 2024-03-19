package channels

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelProperties struct {
	ChannelType                     *ChannelType              `json:"channelType,omitempty"`
	ExpirationTimeIfNotActivatedUtc *string                   `json:"expirationTimeIfNotActivatedUtc,omitempty"`
	MessageForActivation            *string                   `json:"messageForActivation,omitempty"`
	PartnerDestinationInfo          PartnerDestinationInfo    `json:"partnerDestinationInfo"`
	PartnerTopicInfo                *PartnerTopicInfo         `json:"partnerTopicInfo,omitempty"`
	ProvisioningState               *ChannelProvisioningState `json:"provisioningState,omitempty"`
	ReadinessState                  *ReadinessState           `json:"readinessState,omitempty"`
}

func (o *ChannelProperties) GetExpirationTimeIfNotActivatedUtcAsTime() (*time.Time, error) {
	if o.ExpirationTimeIfNotActivatedUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationTimeIfNotActivatedUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *ChannelProperties) SetExpirationTimeIfNotActivatedUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationTimeIfNotActivatedUtc = &formatted
}

var _ json.Unmarshaler = &ChannelProperties{}

func (s *ChannelProperties) UnmarshalJSON(bytes []byte) error {
	type alias ChannelProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ChannelProperties: %+v", err)
	}

	s.ChannelType = decoded.ChannelType
	s.ExpirationTimeIfNotActivatedUtc = decoded.ExpirationTimeIfNotActivatedUtc
	s.MessageForActivation = decoded.MessageForActivation
	s.PartnerTopicInfo = decoded.PartnerTopicInfo
	s.ProvisioningState = decoded.ProvisioningState
	s.ReadinessState = decoded.ReadinessState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChannelProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["partnerDestinationInfo"]; ok {
		impl, err := unmarshalPartnerDestinationInfoImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PartnerDestinationInfo' for 'ChannelProperties': %+v", err)
		}
		s.PartnerDestinationInfo = impl
	}
	return nil
}
