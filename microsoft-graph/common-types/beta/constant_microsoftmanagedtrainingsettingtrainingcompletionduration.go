package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedTrainingSettingTrainingCompletionDuration string

const (
	MicrosoftManagedTrainingSettingTrainingCompletionDuration_Fortnite MicrosoftManagedTrainingSettingTrainingCompletionDuration = "fortnite"
	MicrosoftManagedTrainingSettingTrainingCompletionDuration_Month    MicrosoftManagedTrainingSettingTrainingCompletionDuration = "month"
	MicrosoftManagedTrainingSettingTrainingCompletionDuration_Week     MicrosoftManagedTrainingSettingTrainingCompletionDuration = "week"
)

func PossibleValuesForMicrosoftManagedTrainingSettingTrainingCompletionDuration() []string {
	return []string{
		string(MicrosoftManagedTrainingSettingTrainingCompletionDuration_Fortnite),
		string(MicrosoftManagedTrainingSettingTrainingCompletionDuration_Month),
		string(MicrosoftManagedTrainingSettingTrainingCompletionDuration_Week),
	}
}

func (s *MicrosoftManagedTrainingSettingTrainingCompletionDuration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftManagedTrainingSettingTrainingCompletionDuration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftManagedTrainingSettingTrainingCompletionDuration(input string) (*MicrosoftManagedTrainingSettingTrainingCompletionDuration, error) {
	vals := map[string]MicrosoftManagedTrainingSettingTrainingCompletionDuration{
		"fortnite": MicrosoftManagedTrainingSettingTrainingCompletionDuration_Fortnite,
		"month":    MicrosoftManagedTrainingSettingTrainingCompletionDuration_Month,
		"week":     MicrosoftManagedTrainingSettingTrainingCompletionDuration_Week,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftManagedTrainingSettingTrainingCompletionDuration(input)
	return &out, nil
}
