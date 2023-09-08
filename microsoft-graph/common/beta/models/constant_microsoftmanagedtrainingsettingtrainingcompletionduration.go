package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedTrainingSettingTrainingCompletionDuration string

const (
	MicrosoftManagedTrainingSettingTrainingCompletionDurationfortnite MicrosoftManagedTrainingSettingTrainingCompletionDuration = "Fortnite"
	MicrosoftManagedTrainingSettingTrainingCompletionDurationmonth    MicrosoftManagedTrainingSettingTrainingCompletionDuration = "Month"
	MicrosoftManagedTrainingSettingTrainingCompletionDurationweek     MicrosoftManagedTrainingSettingTrainingCompletionDuration = "Week"
)

func PossibleValuesForMicrosoftManagedTrainingSettingTrainingCompletionDuration() []string {
	return []string{
		string(MicrosoftManagedTrainingSettingTrainingCompletionDurationfortnite),
		string(MicrosoftManagedTrainingSettingTrainingCompletionDurationmonth),
		string(MicrosoftManagedTrainingSettingTrainingCompletionDurationweek),
	}
}

func parseMicrosoftManagedTrainingSettingTrainingCompletionDuration(input string) (*MicrosoftManagedTrainingSettingTrainingCompletionDuration, error) {
	vals := map[string]MicrosoftManagedTrainingSettingTrainingCompletionDuration{
		"fortnite": MicrosoftManagedTrainingSettingTrainingCompletionDurationfortnite,
		"month":    MicrosoftManagedTrainingSettingTrainingCompletionDurationmonth,
		"week":     MicrosoftManagedTrainingSettingTrainingCompletionDurationweek,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftManagedTrainingSettingTrainingCompletionDuration(input)
	return &out, nil
}
