package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyOperationOperationType string

const (
	GroupPolicyOperationOperationType_AddLanguageFiles    GroupPolicyOperationOperationType = "addLanguageFiles"
	GroupPolicyOperationOperationType_None                GroupPolicyOperationOperationType = "none"
	GroupPolicyOperationOperationType_Remove              GroupPolicyOperationOperationType = "remove"
	GroupPolicyOperationOperationType_RemoveLanguageFiles GroupPolicyOperationOperationType = "removeLanguageFiles"
	GroupPolicyOperationOperationType_UpdateLanguageFiles GroupPolicyOperationOperationType = "updateLanguageFiles"
	GroupPolicyOperationOperationType_Upload              GroupPolicyOperationOperationType = "upload"
	GroupPolicyOperationOperationType_UploadNewVersion    GroupPolicyOperationOperationType = "uploadNewVersion"
)

func PossibleValuesForGroupPolicyOperationOperationType() []string {
	return []string{
		string(GroupPolicyOperationOperationType_AddLanguageFiles),
		string(GroupPolicyOperationOperationType_None),
		string(GroupPolicyOperationOperationType_Remove),
		string(GroupPolicyOperationOperationType_RemoveLanguageFiles),
		string(GroupPolicyOperationOperationType_UpdateLanguageFiles),
		string(GroupPolicyOperationOperationType_Upload),
		string(GroupPolicyOperationOperationType_UploadNewVersion),
	}
}

func (s *GroupPolicyOperationOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyOperationOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyOperationOperationType(input string) (*GroupPolicyOperationOperationType, error) {
	vals := map[string]GroupPolicyOperationOperationType{
		"addlanguagefiles":    GroupPolicyOperationOperationType_AddLanguageFiles,
		"none":                GroupPolicyOperationOperationType_None,
		"remove":              GroupPolicyOperationOperationType_Remove,
		"removelanguagefiles": GroupPolicyOperationOperationType_RemoveLanguageFiles,
		"updatelanguagefiles": GroupPolicyOperationOperationType_UpdateLanguageFiles,
		"upload":              GroupPolicyOperationOperationType_Upload,
		"uploadnewversion":    GroupPolicyOperationOperationType_UploadNewVersion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyOperationOperationType(input)
	return &out, nil
}
