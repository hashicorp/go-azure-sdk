package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppPublishingState string

const (
	AndroidManagedStoreAppPublishingState_NotPublished AndroidManagedStoreAppPublishingState = "notPublished"
	AndroidManagedStoreAppPublishingState_Processing   AndroidManagedStoreAppPublishingState = "processing"
	AndroidManagedStoreAppPublishingState_Published    AndroidManagedStoreAppPublishingState = "published"
)

func PossibleValuesForAndroidManagedStoreAppPublishingState() []string {
	return []string{
		string(AndroidManagedStoreAppPublishingState_NotPublished),
		string(AndroidManagedStoreAppPublishingState_Processing),
		string(AndroidManagedStoreAppPublishingState_Published),
	}
}

func (s *AndroidManagedStoreAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAppPublishingState(input string) (*AndroidManagedStoreAppPublishingState, error) {
	vals := map[string]AndroidManagedStoreAppPublishingState{
		"notpublished": AndroidManagedStoreAppPublishingState_NotPublished,
		"processing":   AndroidManagedStoreAppPublishingState_Processing,
		"published":    AndroidManagedStoreAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAppPublishingState(input)
	return &out, nil
}
