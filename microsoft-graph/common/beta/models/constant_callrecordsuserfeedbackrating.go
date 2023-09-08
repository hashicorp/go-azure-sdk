package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsUserFeedbackRating string

const (
	CallRecordsUserFeedbackRatingbad       CallRecordsUserFeedbackRating = "Bad"
	CallRecordsUserFeedbackRatingexcellent CallRecordsUserFeedbackRating = "Excellent"
	CallRecordsUserFeedbackRatingfair      CallRecordsUserFeedbackRating = "Fair"
	CallRecordsUserFeedbackRatinggood      CallRecordsUserFeedbackRating = "Good"
	CallRecordsUserFeedbackRatingnotRated  CallRecordsUserFeedbackRating = "NotRated"
	CallRecordsUserFeedbackRatingpoor      CallRecordsUserFeedbackRating = "Poor"
)

func PossibleValuesForCallRecordsUserFeedbackRating() []string {
	return []string{
		string(CallRecordsUserFeedbackRatingbad),
		string(CallRecordsUserFeedbackRatingexcellent),
		string(CallRecordsUserFeedbackRatingfair),
		string(CallRecordsUserFeedbackRatinggood),
		string(CallRecordsUserFeedbackRatingnotRated),
		string(CallRecordsUserFeedbackRatingpoor),
	}
}

func parseCallRecordsUserFeedbackRating(input string) (*CallRecordsUserFeedbackRating, error) {
	vals := map[string]CallRecordsUserFeedbackRating{
		"bad":       CallRecordsUserFeedbackRatingbad,
		"excellent": CallRecordsUserFeedbackRatingexcellent,
		"fair":      CallRecordsUserFeedbackRatingfair,
		"good":      CallRecordsUserFeedbackRatinggood,
		"notrated":  CallRecordsUserFeedbackRatingnotRated,
		"poor":      CallRecordsUserFeedbackRatingpoor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsUserFeedbackRating(input)
	return &out, nil
}
