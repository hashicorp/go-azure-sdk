package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppPublishingState string

const (
	IosVppAppPublishingState_NotPublished IosVppAppPublishingState = "notPublished"
	IosVppAppPublishingState_Processing   IosVppAppPublishingState = "processing"
	IosVppAppPublishingState_Published    IosVppAppPublishingState = "published"
)

func PossibleValuesForIosVppAppPublishingState() []string {
	return []string{
		string(IosVppAppPublishingState_NotPublished),
		string(IosVppAppPublishingState_Processing),
		string(IosVppAppPublishingState_Published),
	}
}

func (s *IosVppAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVppAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVppAppPublishingState(input string) (*IosVppAppPublishingState, error) {
	vals := map[string]IosVppAppPublishingState{
		"notpublished": IosVppAppPublishingState_NotPublished,
		"processing":   IosVppAppPublishingState_Processing,
		"published":    IosVppAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppAppPublishingState(input)
	return &out, nil
}
