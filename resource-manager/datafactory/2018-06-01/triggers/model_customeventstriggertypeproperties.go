package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomEventsTriggerTypeProperties struct {
	Events            []interface{} `json:"events"`
	Scope             string        `json:"scope"`
	SubjectBeginsWith *string       `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith   *string       `json:"subjectEndsWith,omitempty"`
}
