package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddLargeGalleryViewOperationStatus string

const (
	AddLargeGalleryViewOperationStatusCompleted  AddLargeGalleryViewOperationStatus = "Completed"
	AddLargeGalleryViewOperationStatusFailed     AddLargeGalleryViewOperationStatus = "Failed"
	AddLargeGalleryViewOperationStatusNotStarted AddLargeGalleryViewOperationStatus = "NotStarted"
	AddLargeGalleryViewOperationStatusRunning    AddLargeGalleryViewOperationStatus = "Running"
)

func PossibleValuesForAddLargeGalleryViewOperationStatus() []string {
	return []string{
		string(AddLargeGalleryViewOperationStatusCompleted),
		string(AddLargeGalleryViewOperationStatusFailed),
		string(AddLargeGalleryViewOperationStatusNotStarted),
		string(AddLargeGalleryViewOperationStatusRunning),
	}
}

func parseAddLargeGalleryViewOperationStatus(input string) (*AddLargeGalleryViewOperationStatus, error) {
	vals := map[string]AddLargeGalleryViewOperationStatus{
		"completed":  AddLargeGalleryViewOperationStatusCompleted,
		"failed":     AddLargeGalleryViewOperationStatusFailed,
		"notstarted": AddLargeGalleryViewOperationStatusNotStarted,
		"running":    AddLargeGalleryViewOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddLargeGalleryViewOperationStatus(input)
	return &out, nil
}
