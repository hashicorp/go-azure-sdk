package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftCustomTrainingSettingTrainingCompletionDuration string

const (
	MicrosoftCustomTrainingSettingTrainingCompletionDuration_Fortnite MicrosoftCustomTrainingSettingTrainingCompletionDuration = "fortnite"
	MicrosoftCustomTrainingSettingTrainingCompletionDuration_Month    MicrosoftCustomTrainingSettingTrainingCompletionDuration = "month"
	MicrosoftCustomTrainingSettingTrainingCompletionDuration_Week     MicrosoftCustomTrainingSettingTrainingCompletionDuration = "week"
)

func PossibleValuesForMicrosoftCustomTrainingSettingTrainingCompletionDuration() []string {
	return []string{
		string(MicrosoftCustomTrainingSettingTrainingCompletionDuration_Fortnite),
		string(MicrosoftCustomTrainingSettingTrainingCompletionDuration_Month),
		string(MicrosoftCustomTrainingSettingTrainingCompletionDuration_Week),
	}
}

func (s *MicrosoftCustomTrainingSettingTrainingCompletionDuration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftCustomTrainingSettingTrainingCompletionDuration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftCustomTrainingSettingTrainingCompletionDuration(input string) (*MicrosoftCustomTrainingSettingTrainingCompletionDuration, error) {
	vals := map[string]MicrosoftCustomTrainingSettingTrainingCompletionDuration{
		"fortnite": MicrosoftCustomTrainingSettingTrainingCompletionDuration_Fortnite,
		"month":    MicrosoftCustomTrainingSettingTrainingCompletionDuration_Month,
		"week":     MicrosoftCustomTrainingSettingTrainingCompletionDuration_Week,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftCustomTrainingSettingTrainingCompletionDuration(input)
	return &out, nil
}
