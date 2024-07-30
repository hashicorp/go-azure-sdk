package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogCollectionRequestStatus string

const (
	AppLogCollectionRequestStatus_Completed AppLogCollectionRequestStatus = "completed"
	AppLogCollectionRequestStatus_Failed    AppLogCollectionRequestStatus = "failed"
	AppLogCollectionRequestStatus_Pending   AppLogCollectionRequestStatus = "pending"
)

func PossibleValuesForAppLogCollectionRequestStatus() []string {
	return []string{
		string(AppLogCollectionRequestStatus_Completed),
		string(AppLogCollectionRequestStatus_Failed),
		string(AppLogCollectionRequestStatus_Pending),
	}
}

func (s *AppLogCollectionRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppLogCollectionRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppLogCollectionRequestStatus(input string) (*AppLogCollectionRequestStatus, error) {
	vals := map[string]AppLogCollectionRequestStatus{
		"completed": AppLogCollectionRequestStatus_Completed,
		"failed":    AppLogCollectionRequestStatus_Failed,
		"pending":   AppLogCollectionRequestStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppLogCollectionRequestStatus(input)
	return &out, nil
}
