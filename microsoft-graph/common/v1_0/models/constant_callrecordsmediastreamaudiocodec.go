package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStreamAudioCodec string

const (
	CallRecordsMediaStreamAudioCodecamrWide           CallRecordsMediaStreamAudioCodec = "AmrWide"
	CallRecordsMediaStreamAudioCodeccn                CallRecordsMediaStreamAudioCodec = "Cn"
	CallRecordsMediaStreamAudioCodecg722              CallRecordsMediaStreamAudioCodec = "G722"
	CallRecordsMediaStreamAudioCodecg7221             CallRecordsMediaStreamAudioCodec = "G7221"
	CallRecordsMediaStreamAudioCodecg7221c            CallRecordsMediaStreamAudioCodec = "G7221c"
	CallRecordsMediaStreamAudioCodecg729              CallRecordsMediaStreamAudioCodec = "G729"
	CallRecordsMediaStreamAudioCodecinvalid           CallRecordsMediaStreamAudioCodec = "Invalid"
	CallRecordsMediaStreamAudioCodecmuchv2            CallRecordsMediaStreamAudioCodec = "Muchv2"
	CallRecordsMediaStreamAudioCodecmultiChannelAudio CallRecordsMediaStreamAudioCodec = "MultiChannelAudio"
	CallRecordsMediaStreamAudioCodecopus              CallRecordsMediaStreamAudioCodec = "Opus"
	CallRecordsMediaStreamAudioCodecpcma              CallRecordsMediaStreamAudioCodec = "Pcma"
	CallRecordsMediaStreamAudioCodecpcmu              CallRecordsMediaStreamAudioCodec = "Pcmu"
	CallRecordsMediaStreamAudioCodecrtAudio16         CallRecordsMediaStreamAudioCodec = "RtAudio16"
	CallRecordsMediaStreamAudioCodecrtAudio8          CallRecordsMediaStreamAudioCodec = "RtAudio8"
	CallRecordsMediaStreamAudioCodecsatin             CallRecordsMediaStreamAudioCodec = "Satin"
	CallRecordsMediaStreamAudioCodecsatinFullband     CallRecordsMediaStreamAudioCodec = "SatinFullband"
	CallRecordsMediaStreamAudioCodecsilk              CallRecordsMediaStreamAudioCodec = "Silk"
	CallRecordsMediaStreamAudioCodecsilkNarrow        CallRecordsMediaStreamAudioCodec = "SilkNarrow"
	CallRecordsMediaStreamAudioCodecsilkWide          CallRecordsMediaStreamAudioCodec = "SilkWide"
	CallRecordsMediaStreamAudioCodecsiren             CallRecordsMediaStreamAudioCodec = "Siren"
	CallRecordsMediaStreamAudioCodecunknown           CallRecordsMediaStreamAudioCodec = "Unknown"
	CallRecordsMediaStreamAudioCodecxmsRta            CallRecordsMediaStreamAudioCodec = "XmsRta"
)

func PossibleValuesForCallRecordsMediaStreamAudioCodec() []string {
	return []string{
		string(CallRecordsMediaStreamAudioCodecamrWide),
		string(CallRecordsMediaStreamAudioCodeccn),
		string(CallRecordsMediaStreamAudioCodecg722),
		string(CallRecordsMediaStreamAudioCodecg7221),
		string(CallRecordsMediaStreamAudioCodecg7221c),
		string(CallRecordsMediaStreamAudioCodecg729),
		string(CallRecordsMediaStreamAudioCodecinvalid),
		string(CallRecordsMediaStreamAudioCodecmuchv2),
		string(CallRecordsMediaStreamAudioCodecmultiChannelAudio),
		string(CallRecordsMediaStreamAudioCodecopus),
		string(CallRecordsMediaStreamAudioCodecpcma),
		string(CallRecordsMediaStreamAudioCodecpcmu),
		string(CallRecordsMediaStreamAudioCodecrtAudio16),
		string(CallRecordsMediaStreamAudioCodecrtAudio8),
		string(CallRecordsMediaStreamAudioCodecsatin),
		string(CallRecordsMediaStreamAudioCodecsatinFullband),
		string(CallRecordsMediaStreamAudioCodecsilk),
		string(CallRecordsMediaStreamAudioCodecsilkNarrow),
		string(CallRecordsMediaStreamAudioCodecsilkWide),
		string(CallRecordsMediaStreamAudioCodecsiren),
		string(CallRecordsMediaStreamAudioCodecunknown),
		string(CallRecordsMediaStreamAudioCodecxmsRta),
	}
}

func parseCallRecordsMediaStreamAudioCodec(input string) (*CallRecordsMediaStreamAudioCodec, error) {
	vals := map[string]CallRecordsMediaStreamAudioCodec{
		"amrwide":           CallRecordsMediaStreamAudioCodecamrWide,
		"cn":                CallRecordsMediaStreamAudioCodeccn,
		"g722":              CallRecordsMediaStreamAudioCodecg722,
		"g7221":             CallRecordsMediaStreamAudioCodecg7221,
		"g7221c":            CallRecordsMediaStreamAudioCodecg7221c,
		"g729":              CallRecordsMediaStreamAudioCodecg729,
		"invalid":           CallRecordsMediaStreamAudioCodecinvalid,
		"muchv2":            CallRecordsMediaStreamAudioCodecmuchv2,
		"multichannelaudio": CallRecordsMediaStreamAudioCodecmultiChannelAudio,
		"opus":              CallRecordsMediaStreamAudioCodecopus,
		"pcma":              CallRecordsMediaStreamAudioCodecpcma,
		"pcmu":              CallRecordsMediaStreamAudioCodecpcmu,
		"rtaudio16":         CallRecordsMediaStreamAudioCodecrtAudio16,
		"rtaudio8":          CallRecordsMediaStreamAudioCodecrtAudio8,
		"satin":             CallRecordsMediaStreamAudioCodecsatin,
		"satinfullband":     CallRecordsMediaStreamAudioCodecsatinFullband,
		"silk":              CallRecordsMediaStreamAudioCodecsilk,
		"silknarrow":        CallRecordsMediaStreamAudioCodecsilkNarrow,
		"silkwide":          CallRecordsMediaStreamAudioCodecsilkWide,
		"siren":             CallRecordsMediaStreamAudioCodecsiren,
		"unknown":           CallRecordsMediaStreamAudioCodecunknown,
		"xmsrta":            CallRecordsMediaStreamAudioCodecxmsRta,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsMediaStreamAudioCodec(input)
	return &out, nil
}
