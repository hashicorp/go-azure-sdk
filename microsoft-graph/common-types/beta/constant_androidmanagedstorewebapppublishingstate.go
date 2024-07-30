package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreWebAppPublishingState string

const (
	AndroidManagedStoreWebAppPublishingState_NotPublished AndroidManagedStoreWebAppPublishingState = "notPublished"
	AndroidManagedStoreWebAppPublishingState_Processing   AndroidManagedStoreWebAppPublishingState = "processing"
	AndroidManagedStoreWebAppPublishingState_Published    AndroidManagedStoreWebAppPublishingState = "published"
)

func PossibleValuesForAndroidManagedStoreWebAppPublishingState() []string {
	return []string{
		string(AndroidManagedStoreWebAppPublishingState_NotPublished),
		string(AndroidManagedStoreWebAppPublishingState_Processing),
		string(AndroidManagedStoreWebAppPublishingState_Published),
	}
}

func (s *AndroidManagedStoreWebAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreWebAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreWebAppPublishingState(input string) (*AndroidManagedStoreWebAppPublishingState, error) {
	vals := map[string]AndroidManagedStoreWebAppPublishingState{
		"notpublished": AndroidManagedStoreWebAppPublishingState_NotPublished,
		"processing":   AndroidManagedStoreWebAppPublishingState_Processing,
		"published":    AndroidManagedStoreWebAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreWebAppPublishingState(input)
	return &out, nil
}
