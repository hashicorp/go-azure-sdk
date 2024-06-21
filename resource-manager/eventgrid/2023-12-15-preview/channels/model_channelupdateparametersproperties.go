package channels

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelUpdateParametersProperties struct {
	ExpirationTimeIfNotActivatedUtc *string                      `json:"expirationTimeIfNotActivatedUtc,omitempty"`
	PartnerDestinationInfo          PartnerUpdateDestinationInfo `json:"partnerDestinationInfo"`
	PartnerTopicInfo                *PartnerUpdateTopicInfo      `json:"partnerTopicInfo,omitempty"`
}

var _ json.Unmarshaler = &ChannelUpdateParametersProperties{}

func (s *ChannelUpdateParametersProperties) UnmarshalJSON(bytes []byte) error {
	type alias ChannelUpdateParametersProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ChannelUpdateParametersProperties: %+v", err)
	}

	s.ExpirationTimeIfNotActivatedUtc = decoded.ExpirationTimeIfNotActivatedUtc
	s.PartnerTopicInfo = decoded.PartnerTopicInfo

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChannelUpdateParametersProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["partnerDestinationInfo"]; ok {
		impl, err := unmarshalPartnerUpdateDestinationInfoImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PartnerDestinationInfo' for 'ChannelUpdateParametersProperties': %+v", err)
		}
		s.PartnerDestinationInfo = impl
	}
	return nil
}
