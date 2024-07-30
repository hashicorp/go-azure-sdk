package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStreamAudioCodec string

const (
	CallRecordsMediaStreamAudioCodec_AmrWide           CallRecordsMediaStreamAudioCodec = "amrWide"
	CallRecordsMediaStreamAudioCodec_Cn                CallRecordsMediaStreamAudioCodec = "cn"
	CallRecordsMediaStreamAudioCodec_G722              CallRecordsMediaStreamAudioCodec = "g722"
	CallRecordsMediaStreamAudioCodec_G7221             CallRecordsMediaStreamAudioCodec = "g7221"
	CallRecordsMediaStreamAudioCodec_G7221c            CallRecordsMediaStreamAudioCodec = "g7221c"
	CallRecordsMediaStreamAudioCodec_G729              CallRecordsMediaStreamAudioCodec = "g729"
	CallRecordsMediaStreamAudioCodec_Invalid           CallRecordsMediaStreamAudioCodec = "invalid"
	CallRecordsMediaStreamAudioCodec_Muchv2            CallRecordsMediaStreamAudioCodec = "muchv2"
	CallRecordsMediaStreamAudioCodec_MultiChannelAudio CallRecordsMediaStreamAudioCodec = "multiChannelAudio"
	CallRecordsMediaStreamAudioCodec_Opus              CallRecordsMediaStreamAudioCodec = "opus"
	CallRecordsMediaStreamAudioCodec_Pcma              CallRecordsMediaStreamAudioCodec = "pcma"
	CallRecordsMediaStreamAudioCodec_Pcmu              CallRecordsMediaStreamAudioCodec = "pcmu"
	CallRecordsMediaStreamAudioCodec_RtAudio16         CallRecordsMediaStreamAudioCodec = "rtAudio16"
	CallRecordsMediaStreamAudioCodec_RtAudio8          CallRecordsMediaStreamAudioCodec = "rtAudio8"
	CallRecordsMediaStreamAudioCodec_Satin             CallRecordsMediaStreamAudioCodec = "satin"
	CallRecordsMediaStreamAudioCodec_SatinFullband     CallRecordsMediaStreamAudioCodec = "satinFullband"
	CallRecordsMediaStreamAudioCodec_Silk              CallRecordsMediaStreamAudioCodec = "silk"
	CallRecordsMediaStreamAudioCodec_SilkNarrow        CallRecordsMediaStreamAudioCodec = "silkNarrow"
	CallRecordsMediaStreamAudioCodec_SilkWide          CallRecordsMediaStreamAudioCodec = "silkWide"
	CallRecordsMediaStreamAudioCodec_Siren             CallRecordsMediaStreamAudioCodec = "siren"
	CallRecordsMediaStreamAudioCodec_Unknown           CallRecordsMediaStreamAudioCodec = "unknown"
	CallRecordsMediaStreamAudioCodec_XmsRta            CallRecordsMediaStreamAudioCodec = "xmsRta"
)

func PossibleValuesForCallRecordsMediaStreamAudioCodec() []string {
	return []string{
		string(CallRecordsMediaStreamAudioCodec_AmrWide),
		string(CallRecordsMediaStreamAudioCodec_Cn),
		string(CallRecordsMediaStreamAudioCodec_G722),
		string(CallRecordsMediaStreamAudioCodec_G7221),
		string(CallRecordsMediaStreamAudioCodec_G7221c),
		string(CallRecordsMediaStreamAudioCodec_G729),
		string(CallRecordsMediaStreamAudioCodec_Invalid),
		string(CallRecordsMediaStreamAudioCodec_Muchv2),
		string(CallRecordsMediaStreamAudioCodec_MultiChannelAudio),
		string(CallRecordsMediaStreamAudioCodec_Opus),
		string(CallRecordsMediaStreamAudioCodec_Pcma),
		string(CallRecordsMediaStreamAudioCodec_Pcmu),
		string(CallRecordsMediaStreamAudioCodec_RtAudio16),
		string(CallRecordsMediaStreamAudioCodec_RtAudio8),
		string(CallRecordsMediaStreamAudioCodec_Satin),
		string(CallRecordsMediaStreamAudioCodec_SatinFullband),
		string(CallRecordsMediaStreamAudioCodec_Silk),
		string(CallRecordsMediaStreamAudioCodec_SilkNarrow),
		string(CallRecordsMediaStreamAudioCodec_SilkWide),
		string(CallRecordsMediaStreamAudioCodec_Siren),
		string(CallRecordsMediaStreamAudioCodec_Unknown),
		string(CallRecordsMediaStreamAudioCodec_XmsRta),
	}
}

func (s *CallRecordsMediaStreamAudioCodec) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsMediaStreamAudioCodec(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsMediaStreamAudioCodec(input string) (*CallRecordsMediaStreamAudioCodec, error) {
	vals := map[string]CallRecordsMediaStreamAudioCodec{
		"amrwide":           CallRecordsMediaStreamAudioCodec_AmrWide,
		"cn":                CallRecordsMediaStreamAudioCodec_Cn,
		"g722":              CallRecordsMediaStreamAudioCodec_G722,
		"g7221":             CallRecordsMediaStreamAudioCodec_G7221,
		"g7221c":            CallRecordsMediaStreamAudioCodec_G7221c,
		"g729":              CallRecordsMediaStreamAudioCodec_G729,
		"invalid":           CallRecordsMediaStreamAudioCodec_Invalid,
		"muchv2":            CallRecordsMediaStreamAudioCodec_Muchv2,
		"multichannelaudio": CallRecordsMediaStreamAudioCodec_MultiChannelAudio,
		"opus":              CallRecordsMediaStreamAudioCodec_Opus,
		"pcma":              CallRecordsMediaStreamAudioCodec_Pcma,
		"pcmu":              CallRecordsMediaStreamAudioCodec_Pcmu,
		"rtaudio16":         CallRecordsMediaStreamAudioCodec_RtAudio16,
		"rtaudio8":          CallRecordsMediaStreamAudioCodec_RtAudio8,
		"satin":             CallRecordsMediaStreamAudioCodec_Satin,
		"satinfullband":     CallRecordsMediaStreamAudioCodec_SatinFullband,
		"silk":              CallRecordsMediaStreamAudioCodec_Silk,
		"silknarrow":        CallRecordsMediaStreamAudioCodec_SilkNarrow,
		"silkwide":          CallRecordsMediaStreamAudioCodec_SilkWide,
		"siren":             CallRecordsMediaStreamAudioCodec_Siren,
		"unknown":           CallRecordsMediaStreamAudioCodec_Unknown,
		"xmsrta":            CallRecordsMediaStreamAudioCodec_XmsRta,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsMediaStreamAudioCodec(input)
	return &out, nil
}
