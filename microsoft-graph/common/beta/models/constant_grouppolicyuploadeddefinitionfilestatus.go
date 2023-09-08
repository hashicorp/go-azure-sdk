package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedDefinitionFileStatus string

const (
	GroupPolicyUploadedDefinitionFileStatusassigned          GroupPolicyUploadedDefinitionFileStatus = "Assigned"
	GroupPolicyUploadedDefinitionFileStatusavailable         GroupPolicyUploadedDefinitionFileStatus = "Available"
	GroupPolicyUploadedDefinitionFileStatusnone              GroupPolicyUploadedDefinitionFileStatus = "None"
	GroupPolicyUploadedDefinitionFileStatusremovalFailed     GroupPolicyUploadedDefinitionFileStatus = "RemovalFailed"
	GroupPolicyUploadedDefinitionFileStatusremovalInProgress GroupPolicyUploadedDefinitionFileStatus = "RemovalInProgress"
	GroupPolicyUploadedDefinitionFileStatusuploadFailed      GroupPolicyUploadedDefinitionFileStatus = "UploadFailed"
	GroupPolicyUploadedDefinitionFileStatusuploadInProgress  GroupPolicyUploadedDefinitionFileStatus = "UploadInProgress"
)

func PossibleValuesForGroupPolicyUploadedDefinitionFileStatus() []string {
	return []string{
		string(GroupPolicyUploadedDefinitionFileStatusassigned),
		string(GroupPolicyUploadedDefinitionFileStatusavailable),
		string(GroupPolicyUploadedDefinitionFileStatusnone),
		string(GroupPolicyUploadedDefinitionFileStatusremovalFailed),
		string(GroupPolicyUploadedDefinitionFileStatusremovalInProgress),
		string(GroupPolicyUploadedDefinitionFileStatusuploadFailed),
		string(GroupPolicyUploadedDefinitionFileStatusuploadInProgress),
	}
}

func parseGroupPolicyUploadedDefinitionFileStatus(input string) (*GroupPolicyUploadedDefinitionFileStatus, error) {
	vals := map[string]GroupPolicyUploadedDefinitionFileStatus{
		"assigned":          GroupPolicyUploadedDefinitionFileStatusassigned,
		"available":         GroupPolicyUploadedDefinitionFileStatusavailable,
		"none":              GroupPolicyUploadedDefinitionFileStatusnone,
		"removalfailed":     GroupPolicyUploadedDefinitionFileStatusremovalFailed,
		"removalinprogress": GroupPolicyUploadedDefinitionFileStatusremovalInProgress,
		"uploadfailed":      GroupPolicyUploadedDefinitionFileStatusuploadFailed,
		"uploadinprogress":  GroupPolicyUploadedDefinitionFileStatusuploadInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyUploadedDefinitionFileStatus(input)
	return &out, nil
}
