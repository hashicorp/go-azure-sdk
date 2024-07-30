package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddLargeGalleryViewOperationStatus string

const (
	AddLargeGalleryViewOperationStatus_Completed  AddLargeGalleryViewOperationStatus = "Completed"
	AddLargeGalleryViewOperationStatus_Failed     AddLargeGalleryViewOperationStatus = "Failed"
	AddLargeGalleryViewOperationStatus_NotStarted AddLargeGalleryViewOperationStatus = "NotStarted"
	AddLargeGalleryViewOperationStatus_Running    AddLargeGalleryViewOperationStatus = "Running"
)

func PossibleValuesForAddLargeGalleryViewOperationStatus() []string {
	return []string{
		string(AddLargeGalleryViewOperationStatus_Completed),
		string(AddLargeGalleryViewOperationStatus_Failed),
		string(AddLargeGalleryViewOperationStatus_NotStarted),
		string(AddLargeGalleryViewOperationStatus_Running),
	}
}

func (s *AddLargeGalleryViewOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddLargeGalleryViewOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddLargeGalleryViewOperationStatus(input string) (*AddLargeGalleryViewOperationStatus, error) {
	vals := map[string]AddLargeGalleryViewOperationStatus{
		"completed":  AddLargeGalleryViewOperationStatus_Completed,
		"failed":     AddLargeGalleryViewOperationStatus_Failed,
		"notstarted": AddLargeGalleryViewOperationStatus_NotStarted,
		"running":    AddLargeGalleryViewOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddLargeGalleryViewOperationStatus(input)
	return &out, nil
}
