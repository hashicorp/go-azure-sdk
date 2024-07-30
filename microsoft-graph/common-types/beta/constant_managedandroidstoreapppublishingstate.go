package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidStoreAppPublishingState string

const (
	ManagedAndroidStoreAppPublishingState_NotPublished ManagedAndroidStoreAppPublishingState = "notPublished"
	ManagedAndroidStoreAppPublishingState_Processing   ManagedAndroidStoreAppPublishingState = "processing"
	ManagedAndroidStoreAppPublishingState_Published    ManagedAndroidStoreAppPublishingState = "published"
)

func PossibleValuesForManagedAndroidStoreAppPublishingState() []string {
	return []string{
		string(ManagedAndroidStoreAppPublishingState_NotPublished),
		string(ManagedAndroidStoreAppPublishingState_Processing),
		string(ManagedAndroidStoreAppPublishingState_Published),
	}
}

func (s *ManagedAndroidStoreAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAndroidStoreAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAndroidStoreAppPublishingState(input string) (*ManagedAndroidStoreAppPublishingState, error) {
	vals := map[string]ManagedAndroidStoreAppPublishingState{
		"notpublished": ManagedAndroidStoreAppPublishingState_NotPublished,
		"processing":   ManagedAndroidStoreAppPublishingState_Processing,
		"published":    ManagedAndroidStoreAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidStoreAppPublishingState(input)
	return &out, nil
}
