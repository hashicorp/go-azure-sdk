package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlterationResponseQueryAlterationType string

const (
	AlterationResponseQueryAlterationType_Modification AlterationResponseQueryAlterationType = "modification"
	AlterationResponseQueryAlterationType_Suggestion   AlterationResponseQueryAlterationType = "suggestion"
)

func PossibleValuesForAlterationResponseQueryAlterationType() []string {
	return []string{
		string(AlterationResponseQueryAlterationType_Modification),
		string(AlterationResponseQueryAlterationType_Suggestion),
	}
}

func (s *AlterationResponseQueryAlterationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlterationResponseQueryAlterationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlterationResponseQueryAlterationType(input string) (*AlterationResponseQueryAlterationType, error) {
	vals := map[string]AlterationResponseQueryAlterationType{
		"modification": AlterationResponseQueryAlterationType_Modification,
		"suggestion":   AlterationResponseQueryAlterationType_Suggestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlterationResponseQueryAlterationType(input)
	return &out, nil
}
