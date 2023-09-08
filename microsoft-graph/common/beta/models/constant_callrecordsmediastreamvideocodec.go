package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStreamVideoCodec string

const (
	CallRecordsMediaStreamVideoCodecav1     CallRecordsMediaStreamVideoCodec = "Av1"
	CallRecordsMediaStreamVideoCodech263    CallRecordsMediaStreamVideoCodec = "H263"
	CallRecordsMediaStreamVideoCodech264    CallRecordsMediaStreamVideoCodec = "H264"
	CallRecordsMediaStreamVideoCodech264s   CallRecordsMediaStreamVideoCodec = "H264s"
	CallRecordsMediaStreamVideoCodech264uc  CallRecordsMediaStreamVideoCodec = "H264uc"
	CallRecordsMediaStreamVideoCodech265    CallRecordsMediaStreamVideoCodec = "H265"
	CallRecordsMediaStreamVideoCodecinvalid CallRecordsMediaStreamVideoCodec = "Invalid"
	CallRecordsMediaStreamVideoCodecrtVideo CallRecordsMediaStreamVideoCodec = "RtVideo"
	CallRecordsMediaStreamVideoCodecrtvc1   CallRecordsMediaStreamVideoCodec = "Rtvc1"
	CallRecordsMediaStreamVideoCodecunknown CallRecordsMediaStreamVideoCodec = "Unknown"
	CallRecordsMediaStreamVideoCodecxrtvc1  CallRecordsMediaStreamVideoCodec = "Xrtvc1"
)

func PossibleValuesForCallRecordsMediaStreamVideoCodec() []string {
	return []string{
		string(CallRecordsMediaStreamVideoCodecav1),
		string(CallRecordsMediaStreamVideoCodech263),
		string(CallRecordsMediaStreamVideoCodech264),
		string(CallRecordsMediaStreamVideoCodech264s),
		string(CallRecordsMediaStreamVideoCodech264uc),
		string(CallRecordsMediaStreamVideoCodech265),
		string(CallRecordsMediaStreamVideoCodecinvalid),
		string(CallRecordsMediaStreamVideoCodecrtVideo),
		string(CallRecordsMediaStreamVideoCodecrtvc1),
		string(CallRecordsMediaStreamVideoCodecunknown),
		string(CallRecordsMediaStreamVideoCodecxrtvc1),
	}
}

func parseCallRecordsMediaStreamVideoCodec(input string) (*CallRecordsMediaStreamVideoCodec, error) {
	vals := map[string]CallRecordsMediaStreamVideoCodec{
		"av1":     CallRecordsMediaStreamVideoCodecav1,
		"h263":    CallRecordsMediaStreamVideoCodech263,
		"h264":    CallRecordsMediaStreamVideoCodech264,
		"h264s":   CallRecordsMediaStreamVideoCodech264s,
		"h264uc":  CallRecordsMediaStreamVideoCodech264uc,
		"h265":    CallRecordsMediaStreamVideoCodech265,
		"invalid": CallRecordsMediaStreamVideoCodecinvalid,
		"rtvideo": CallRecordsMediaStreamVideoCodecrtVideo,
		"rtvc1":   CallRecordsMediaStreamVideoCodecrtvc1,
		"unknown": CallRecordsMediaStreamVideoCodecunknown,
		"xrtvc1":  CallRecordsMediaStreamVideoCodecxrtvc1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsMediaStreamVideoCodec(input)
	return &out, nil
}
