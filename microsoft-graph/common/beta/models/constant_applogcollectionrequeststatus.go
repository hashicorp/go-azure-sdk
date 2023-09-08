package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogCollectionRequestStatus string

const (
	AppLogCollectionRequestStatuscompleted AppLogCollectionRequestStatus = "Completed"
	AppLogCollectionRequestStatusfailed    AppLogCollectionRequestStatus = "Failed"
	AppLogCollectionRequestStatuspending   AppLogCollectionRequestStatus = "Pending"
)

func PossibleValuesForAppLogCollectionRequestStatus() []string {
	return []string{
		string(AppLogCollectionRequestStatuscompleted),
		string(AppLogCollectionRequestStatusfailed),
		string(AppLogCollectionRequestStatuspending),
	}
}

func parseAppLogCollectionRequestStatus(input string) (*AppLogCollectionRequestStatus, error) {
	vals := map[string]AppLogCollectionRequestStatus{
		"completed": AppLogCollectionRequestStatuscompleted,
		"failed":    AppLogCollectionRequestStatusfailed,
		"pending":   AppLogCollectionRequestStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppLogCollectionRequestStatus(input)
	return &out, nil
}
