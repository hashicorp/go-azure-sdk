package liveevents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AsyncOperationStatus string

const (
	AsyncOperationStatusFailed     AsyncOperationStatus = "Failed"
	AsyncOperationStatusInProgress AsyncOperationStatus = "InProgress"
	AsyncOperationStatusSucceeded  AsyncOperationStatus = "Succeeded"
)

func PossibleValuesForAsyncOperationStatus() []string {
	return []string{
		string(AsyncOperationStatusFailed),
		string(AsyncOperationStatusInProgress),
		string(AsyncOperationStatusSucceeded),
	}
}

func (s *AsyncOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAsyncOperationStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AsyncOperationStatus(decoded)
	return nil
}

type LiveEventEncodingType string

const (
	LiveEventEncodingTypeNone                     LiveEventEncodingType = "None"
	LiveEventEncodingTypePassthroughBasic         LiveEventEncodingType = "PassthroughBasic"
	LiveEventEncodingTypePassthroughStandard      LiveEventEncodingType = "PassthroughStandard"
	LiveEventEncodingTypePremiumOneZeroEightZerop LiveEventEncodingType = "Premium1080p"
	LiveEventEncodingTypeStandard                 LiveEventEncodingType = "Standard"
)

func PossibleValuesForLiveEventEncodingType() []string {
	return []string{
		string(LiveEventEncodingTypeNone),
		string(LiveEventEncodingTypePassthroughBasic),
		string(LiveEventEncodingTypePassthroughStandard),
		string(LiveEventEncodingTypePremiumOneZeroEightZerop),
		string(LiveEventEncodingTypeStandard),
	}
}

func (s *LiveEventEncodingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForLiveEventEncodingType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = LiveEventEncodingType(decoded)
	return nil
}

type LiveEventInputProtocol string

const (
	LiveEventInputProtocolFragmentedMPFour LiveEventInputProtocol = "FragmentedMP4"
	LiveEventInputProtocolRTMP             LiveEventInputProtocol = "RTMP"
)

func PossibleValuesForLiveEventInputProtocol() []string {
	return []string{
		string(LiveEventInputProtocolFragmentedMPFour),
		string(LiveEventInputProtocolRTMP),
	}
}

func (s *LiveEventInputProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForLiveEventInputProtocol() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = LiveEventInputProtocol(decoded)
	return nil
}

type LiveEventResourceState string

const (
	LiveEventResourceStateAllocating LiveEventResourceState = "Allocating"
	LiveEventResourceStateDeleting   LiveEventResourceState = "Deleting"
	LiveEventResourceStateRunning    LiveEventResourceState = "Running"
	LiveEventResourceStateStandBy    LiveEventResourceState = "StandBy"
	LiveEventResourceStateStarting   LiveEventResourceState = "Starting"
	LiveEventResourceStateStopped    LiveEventResourceState = "Stopped"
	LiveEventResourceStateStopping   LiveEventResourceState = "Stopping"
)

func PossibleValuesForLiveEventResourceState() []string {
	return []string{
		string(LiveEventResourceStateAllocating),
		string(LiveEventResourceStateDeleting),
		string(LiveEventResourceStateRunning),
		string(LiveEventResourceStateStandBy),
		string(LiveEventResourceStateStarting),
		string(LiveEventResourceStateStopped),
		string(LiveEventResourceStateStopping),
	}
}

func (s *LiveEventResourceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForLiveEventResourceState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = LiveEventResourceState(decoded)
	return nil
}

type StreamOptionsFlag string

const (
	StreamOptionsFlagDefault        StreamOptionsFlag = "Default"
	StreamOptionsFlagLowLatency     StreamOptionsFlag = "LowLatency"
	StreamOptionsFlagLowLatencyVTwo StreamOptionsFlag = "LowLatencyV2"
)

func PossibleValuesForStreamOptionsFlag() []string {
	return []string{
		string(StreamOptionsFlagDefault),
		string(StreamOptionsFlagLowLatency),
		string(StreamOptionsFlagLowLatencyVTwo),
	}
}

func (s *StreamOptionsFlag) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStreamOptionsFlag() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StreamOptionsFlag(decoded)
	return nil
}

type StretchMode string

const (
	StretchModeAutoFit  StretchMode = "AutoFit"
	StretchModeAutoSize StretchMode = "AutoSize"
	StretchModeNone     StretchMode = "None"
)

func PossibleValuesForStretchMode() []string {
	return []string{
		string(StretchModeAutoFit),
		string(StretchModeAutoSize),
		string(StretchModeNone),
	}
}

func (s *StretchMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStretchMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StretchMode(decoded)
	return nil
}
