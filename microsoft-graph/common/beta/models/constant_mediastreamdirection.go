package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaStreamDirection string

const (
	MediaStreamDirectioninactive    MediaStreamDirection = "Inactive"
	MediaStreamDirectionreceiveOnly MediaStreamDirection = "ReceiveOnly"
	MediaStreamDirectionsendOnly    MediaStreamDirection = "SendOnly"
	MediaStreamDirectionsendReceive MediaStreamDirection = "SendReceive"
)

func PossibleValuesForMediaStreamDirection() []string {
	return []string{
		string(MediaStreamDirectioninactive),
		string(MediaStreamDirectionreceiveOnly),
		string(MediaStreamDirectionsendOnly),
		string(MediaStreamDirectionsendReceive),
	}
}

func parseMediaStreamDirection(input string) (*MediaStreamDirection, error) {
	vals := map[string]MediaStreamDirection{
		"inactive":    MediaStreamDirectioninactive,
		"receiveonly": MediaStreamDirectionreceiveOnly,
		"sendonly":    MediaStreamDirectionsendOnly,
		"sendreceive": MediaStreamDirectionsendReceive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaStreamDirection(input)
	return &out, nil
}
