package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalActivityResultType string

const (
	ExternalConnectorsExternalActivityResultType_Commented ExternalConnectorsExternalActivityResultType = "commented"
	ExternalConnectorsExternalActivityResultType_Created   ExternalConnectorsExternalActivityResultType = "created"
	ExternalConnectorsExternalActivityResultType_Modified  ExternalConnectorsExternalActivityResultType = "modified"
	ExternalConnectorsExternalActivityResultType_Viewed    ExternalConnectorsExternalActivityResultType = "viewed"
)

func PossibleValuesForExternalConnectorsExternalActivityResultType() []string {
	return []string{
		string(ExternalConnectorsExternalActivityResultType_Commented),
		string(ExternalConnectorsExternalActivityResultType_Created),
		string(ExternalConnectorsExternalActivityResultType_Modified),
		string(ExternalConnectorsExternalActivityResultType_Viewed),
	}
}

func (s *ExternalConnectorsExternalActivityResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsExternalActivityResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsExternalActivityResultType(input string) (*ExternalConnectorsExternalActivityResultType, error) {
	vals := map[string]ExternalConnectorsExternalActivityResultType{
		"commented": ExternalConnectorsExternalActivityResultType_Commented,
		"created":   ExternalConnectorsExternalActivityResultType_Created,
		"modified":  ExternalConnectorsExternalActivityResultType_Modified,
		"viewed":    ExternalConnectorsExternalActivityResultType_Viewed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalActivityResultType(input)
	return &out, nil
}
