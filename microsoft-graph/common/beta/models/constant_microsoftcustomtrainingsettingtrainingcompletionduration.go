package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftCustomTrainingSettingTrainingCompletionDuration string

const (
	MicrosoftCustomTrainingSettingTrainingCompletionDurationfortnite MicrosoftCustomTrainingSettingTrainingCompletionDuration = "Fortnite"
	MicrosoftCustomTrainingSettingTrainingCompletionDurationmonth    MicrosoftCustomTrainingSettingTrainingCompletionDuration = "Month"
	MicrosoftCustomTrainingSettingTrainingCompletionDurationweek     MicrosoftCustomTrainingSettingTrainingCompletionDuration = "Week"
)

func PossibleValuesForMicrosoftCustomTrainingSettingTrainingCompletionDuration() []string {
	return []string{
		string(MicrosoftCustomTrainingSettingTrainingCompletionDurationfortnite),
		string(MicrosoftCustomTrainingSettingTrainingCompletionDurationmonth),
		string(MicrosoftCustomTrainingSettingTrainingCompletionDurationweek),
	}
}

func parseMicrosoftCustomTrainingSettingTrainingCompletionDuration(input string) (*MicrosoftCustomTrainingSettingTrainingCompletionDuration, error) {
	vals := map[string]MicrosoftCustomTrainingSettingTrainingCompletionDuration{
		"fortnite": MicrosoftCustomTrainingSettingTrainingCompletionDurationfortnite,
		"month":    MicrosoftCustomTrainingSettingTrainingCompletionDurationmonth,
		"week":     MicrosoftCustomTrainingSettingTrainingCompletionDurationweek,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftCustomTrainingSettingTrainingCompletionDuration(input)
	return &out, nil
}
