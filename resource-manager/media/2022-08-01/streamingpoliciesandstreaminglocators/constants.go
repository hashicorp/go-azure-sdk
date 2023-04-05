package streamingpoliciesandstreaminglocators

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionScheme string

const (
	EncryptionSchemeCommonEncryptionCbcs EncryptionScheme = "CommonEncryptionCbcs"
	EncryptionSchemeCommonEncryptionCenc EncryptionScheme = "CommonEncryptionCenc"
	EncryptionSchemeEnvelopeEncryption   EncryptionScheme = "EnvelopeEncryption"
	EncryptionSchemeNoEncryption         EncryptionScheme = "NoEncryption"
)

func PossibleValuesForEncryptionScheme() []string {
	return []string{
		string(EncryptionSchemeCommonEncryptionCbcs),
		string(EncryptionSchemeCommonEncryptionCenc),
		string(EncryptionSchemeEnvelopeEncryption),
		string(EncryptionSchemeNoEncryption),
	}
}

func (s *EncryptionScheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForEncryptionScheme() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = EncryptionScheme(decoded)
	return nil
}

type StreamingLocatorContentKeyType string

const (
	StreamingLocatorContentKeyTypeCommonEncryptionCbcs StreamingLocatorContentKeyType = "CommonEncryptionCbcs"
	StreamingLocatorContentKeyTypeCommonEncryptionCenc StreamingLocatorContentKeyType = "CommonEncryptionCenc"
	StreamingLocatorContentKeyTypeEnvelopeEncryption   StreamingLocatorContentKeyType = "EnvelopeEncryption"
)

func PossibleValuesForStreamingLocatorContentKeyType() []string {
	return []string{
		string(StreamingLocatorContentKeyTypeCommonEncryptionCbcs),
		string(StreamingLocatorContentKeyTypeCommonEncryptionCenc),
		string(StreamingLocatorContentKeyTypeEnvelopeEncryption),
	}
}

func (s *StreamingLocatorContentKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStreamingLocatorContentKeyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StreamingLocatorContentKeyType(decoded)
	return nil
}

type StreamingPolicyStreamingProtocol string

const (
	StreamingPolicyStreamingProtocolDash            StreamingPolicyStreamingProtocol = "Dash"
	StreamingPolicyStreamingProtocolDownload        StreamingPolicyStreamingProtocol = "Download"
	StreamingPolicyStreamingProtocolHls             StreamingPolicyStreamingProtocol = "Hls"
	StreamingPolicyStreamingProtocolSmoothStreaming StreamingPolicyStreamingProtocol = "SmoothStreaming"
)

func PossibleValuesForStreamingPolicyStreamingProtocol() []string {
	return []string{
		string(StreamingPolicyStreamingProtocolDash),
		string(StreamingPolicyStreamingProtocolDownload),
		string(StreamingPolicyStreamingProtocolHls),
		string(StreamingPolicyStreamingProtocolSmoothStreaming),
	}
}

func (s *StreamingPolicyStreamingProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForStreamingPolicyStreamingProtocol() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = StreamingPolicyStreamingProtocol(decoded)
	return nil
}

type TrackPropertyCompareOperation string

const (
	TrackPropertyCompareOperationEqual   TrackPropertyCompareOperation = "Equal"
	TrackPropertyCompareOperationUnknown TrackPropertyCompareOperation = "Unknown"
)

func PossibleValuesForTrackPropertyCompareOperation() []string {
	return []string{
		string(TrackPropertyCompareOperationEqual),
		string(TrackPropertyCompareOperationUnknown),
	}
}

func (s *TrackPropertyCompareOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTrackPropertyCompareOperation() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TrackPropertyCompareOperation(decoded)
	return nil
}

type TrackPropertyType string

const (
	TrackPropertyTypeFourCC  TrackPropertyType = "FourCC"
	TrackPropertyTypeUnknown TrackPropertyType = "Unknown"
)

func PossibleValuesForTrackPropertyType() []string {
	return []string{
		string(TrackPropertyTypeFourCC),
		string(TrackPropertyTypeUnknown),
	}
}

func (s *TrackPropertyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTrackPropertyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TrackPropertyType(decoded)
	return nil
}
