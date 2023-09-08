package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyOperationOperationType string

const (
	GroupPolicyOperationOperationTypeaddLanguageFiles    GroupPolicyOperationOperationType = "AddLanguageFiles"
	GroupPolicyOperationOperationTypenone                GroupPolicyOperationOperationType = "None"
	GroupPolicyOperationOperationTyperemove              GroupPolicyOperationOperationType = "Remove"
	GroupPolicyOperationOperationTyperemoveLanguageFiles GroupPolicyOperationOperationType = "RemoveLanguageFiles"
	GroupPolicyOperationOperationTypeupdateLanguageFiles GroupPolicyOperationOperationType = "UpdateLanguageFiles"
	GroupPolicyOperationOperationTypeupload              GroupPolicyOperationOperationType = "Upload"
	GroupPolicyOperationOperationTypeuploadNewVersion    GroupPolicyOperationOperationType = "UploadNewVersion"
)

func PossibleValuesForGroupPolicyOperationOperationType() []string {
	return []string{
		string(GroupPolicyOperationOperationTypeaddLanguageFiles),
		string(GroupPolicyOperationOperationTypenone),
		string(GroupPolicyOperationOperationTyperemove),
		string(GroupPolicyOperationOperationTyperemoveLanguageFiles),
		string(GroupPolicyOperationOperationTypeupdateLanguageFiles),
		string(GroupPolicyOperationOperationTypeupload),
		string(GroupPolicyOperationOperationTypeuploadNewVersion),
	}
}

func parseGroupPolicyOperationOperationType(input string) (*GroupPolicyOperationOperationType, error) {
	vals := map[string]GroupPolicyOperationOperationType{
		"addlanguagefiles":    GroupPolicyOperationOperationTypeaddLanguageFiles,
		"none":                GroupPolicyOperationOperationTypenone,
		"remove":              GroupPolicyOperationOperationTyperemove,
		"removelanguagefiles": GroupPolicyOperationOperationTyperemoveLanguageFiles,
		"updatelanguagefiles": GroupPolicyOperationOperationTypeupdateLanguageFiles,
		"upload":              GroupPolicyOperationOperationTypeupload,
		"uploadnewversion":    GroupPolicyOperationOperationTypeuploadNewVersion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyOperationOperationType(input)
	return &out, nil
}
