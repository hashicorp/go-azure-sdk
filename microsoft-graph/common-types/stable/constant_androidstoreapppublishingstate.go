package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidStoreAppPublishingState string

const (
	AndroidStoreAppPublishingState_NotPublished AndroidStoreAppPublishingState = "notPublished"
	AndroidStoreAppPublishingState_Processing   AndroidStoreAppPublishingState = "processing"
	AndroidStoreAppPublishingState_Published    AndroidStoreAppPublishingState = "published"
)

func PossibleValuesForAndroidStoreAppPublishingState() []string {
	return []string{
		string(AndroidStoreAppPublishingState_NotPublished),
		string(AndroidStoreAppPublishingState_Processing),
		string(AndroidStoreAppPublishingState_Published),
	}
}

func (s *AndroidStoreAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidStoreAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidStoreAppPublishingState(input string) (*AndroidStoreAppPublishingState, error) {
	vals := map[string]AndroidStoreAppPublishingState{
		"notpublished": AndroidStoreAppPublishingState_NotPublished,
		"processing":   AndroidStoreAppPublishingState_Processing,
		"published":    AndroidStoreAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidStoreAppPublishingState(input)
	return &out, nil
}
