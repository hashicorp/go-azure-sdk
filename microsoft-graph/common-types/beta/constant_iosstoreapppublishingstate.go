package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosStoreAppPublishingState string

const (
	IosStoreAppPublishingState_NotPublished IosStoreAppPublishingState = "notPublished"
	IosStoreAppPublishingState_Processing   IosStoreAppPublishingState = "processing"
	IosStoreAppPublishingState_Published    IosStoreAppPublishingState = "published"
)

func PossibleValuesForIosStoreAppPublishingState() []string {
	return []string{
		string(IosStoreAppPublishingState_NotPublished),
		string(IosStoreAppPublishingState_Processing),
		string(IosStoreAppPublishingState_Published),
	}
}

func (s *IosStoreAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosStoreAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosStoreAppPublishingState(input string) (*IosStoreAppPublishingState, error) {
	vals := map[string]IosStoreAppPublishingState{
		"notpublished": IosStoreAppPublishingState_NotPublished,
		"processing":   IosStoreAppPublishingState_Processing,
		"published":    IosStoreAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosStoreAppPublishingState(input)
	return &out, nil
}
