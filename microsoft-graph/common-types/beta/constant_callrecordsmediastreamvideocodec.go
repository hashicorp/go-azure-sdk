package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStreamVideoCodec string

const (
	CallRecordsMediaStreamVideoCodec_Av1     CallRecordsMediaStreamVideoCodec = "av1"
	CallRecordsMediaStreamVideoCodec_H263    CallRecordsMediaStreamVideoCodec = "h263"
	CallRecordsMediaStreamVideoCodec_H264    CallRecordsMediaStreamVideoCodec = "h264"
	CallRecordsMediaStreamVideoCodec_H264s   CallRecordsMediaStreamVideoCodec = "h264s"
	CallRecordsMediaStreamVideoCodec_H264uc  CallRecordsMediaStreamVideoCodec = "h264uc"
	CallRecordsMediaStreamVideoCodec_H265    CallRecordsMediaStreamVideoCodec = "h265"
	CallRecordsMediaStreamVideoCodec_Invalid CallRecordsMediaStreamVideoCodec = "invalid"
	CallRecordsMediaStreamVideoCodec_RtVideo CallRecordsMediaStreamVideoCodec = "rtVideo"
	CallRecordsMediaStreamVideoCodec_Rtvc1   CallRecordsMediaStreamVideoCodec = "rtvc1"
	CallRecordsMediaStreamVideoCodec_Unknown CallRecordsMediaStreamVideoCodec = "unknown"
	CallRecordsMediaStreamVideoCodec_Xrtvc1  CallRecordsMediaStreamVideoCodec = "xrtvc1"
)

func PossibleValuesForCallRecordsMediaStreamVideoCodec() []string {
	return []string{
		string(CallRecordsMediaStreamVideoCodec_Av1),
		string(CallRecordsMediaStreamVideoCodec_H263),
		string(CallRecordsMediaStreamVideoCodec_H264),
		string(CallRecordsMediaStreamVideoCodec_H264s),
		string(CallRecordsMediaStreamVideoCodec_H264uc),
		string(CallRecordsMediaStreamVideoCodec_H265),
		string(CallRecordsMediaStreamVideoCodec_Invalid),
		string(CallRecordsMediaStreamVideoCodec_RtVideo),
		string(CallRecordsMediaStreamVideoCodec_Rtvc1),
		string(CallRecordsMediaStreamVideoCodec_Unknown),
		string(CallRecordsMediaStreamVideoCodec_Xrtvc1),
	}
}

func (s *CallRecordsMediaStreamVideoCodec) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsMediaStreamVideoCodec(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsMediaStreamVideoCodec(input string) (*CallRecordsMediaStreamVideoCodec, error) {
	vals := map[string]CallRecordsMediaStreamVideoCodec{
		"av1":     CallRecordsMediaStreamVideoCodec_Av1,
		"h263":    CallRecordsMediaStreamVideoCodec_H263,
		"h264":    CallRecordsMediaStreamVideoCodec_H264,
		"h264s":   CallRecordsMediaStreamVideoCodec_H264s,
		"h264uc":  CallRecordsMediaStreamVideoCodec_H264uc,
		"h265":    CallRecordsMediaStreamVideoCodec_H265,
		"invalid": CallRecordsMediaStreamVideoCodec_Invalid,
		"rtvideo": CallRecordsMediaStreamVideoCodec_RtVideo,
		"rtvc1":   CallRecordsMediaStreamVideoCodec_Rtvc1,
		"unknown": CallRecordsMediaStreamVideoCodec_Unknown,
		"xrtvc1":  CallRecordsMediaStreamVideoCodec_Xrtvc1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsMediaStreamVideoCodec(input)
	return &out, nil
}
