package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaStreamDirection string

const (
	MediaStreamDirection_Inactive    MediaStreamDirection = "inactive"
	MediaStreamDirection_ReceiveOnly MediaStreamDirection = "receiveOnly"
	MediaStreamDirection_SendOnly    MediaStreamDirection = "sendOnly"
	MediaStreamDirection_SendReceive MediaStreamDirection = "sendReceive"
)

func PossibleValuesForMediaStreamDirection() []string {
	return []string{
		string(MediaStreamDirection_Inactive),
		string(MediaStreamDirection_ReceiveOnly),
		string(MediaStreamDirection_SendOnly),
		string(MediaStreamDirection_SendReceive),
	}
}

func (s *MediaStreamDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaStreamDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaStreamDirection(input string) (*MediaStreamDirection, error) {
	vals := map[string]MediaStreamDirection{
		"inactive":    MediaStreamDirection_Inactive,
		"receiveonly": MediaStreamDirection_ReceiveOnly,
		"sendonly":    MediaStreamDirection_SendOnly,
		"sendreceive": MediaStreamDirection_SendReceive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaStreamDirection(input)
	return &out, nil
}
