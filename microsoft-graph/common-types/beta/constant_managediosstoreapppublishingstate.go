package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIOSStoreAppPublishingState string

const (
	ManagedIOSStoreAppPublishingState_NotPublished ManagedIOSStoreAppPublishingState = "notPublished"
	ManagedIOSStoreAppPublishingState_Processing   ManagedIOSStoreAppPublishingState = "processing"
	ManagedIOSStoreAppPublishingState_Published    ManagedIOSStoreAppPublishingState = "published"
)

func PossibleValuesForManagedIOSStoreAppPublishingState() []string {
	return []string{
		string(ManagedIOSStoreAppPublishingState_NotPublished),
		string(ManagedIOSStoreAppPublishingState_Processing),
		string(ManagedIOSStoreAppPublishingState_Published),
	}
}

func (s *ManagedIOSStoreAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedIOSStoreAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedIOSStoreAppPublishingState(input string) (*ManagedIOSStoreAppPublishingState, error) {
	vals := map[string]ManagedIOSStoreAppPublishingState{
		"notpublished": ManagedIOSStoreAppPublishingState_NotPublished,
		"processing":   ManagedIOSStoreAppPublishingState_Processing,
		"published":    ManagedIOSStoreAppPublishingState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIOSStoreAppPublishingState(input)
	return &out, nil
}
