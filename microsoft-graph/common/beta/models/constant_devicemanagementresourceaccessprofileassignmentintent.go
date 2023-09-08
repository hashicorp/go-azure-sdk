package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceAccessProfileAssignmentIntent string

const (
	DeviceManagementResourceAccessProfileAssignmentIntentapply  DeviceManagementResourceAccessProfileAssignmentIntent = "Apply"
	DeviceManagementResourceAccessProfileAssignmentIntentremove DeviceManagementResourceAccessProfileAssignmentIntent = "Remove"
)

func PossibleValuesForDeviceManagementResourceAccessProfileAssignmentIntent() []string {
	return []string{
		string(DeviceManagementResourceAccessProfileAssignmentIntentapply),
		string(DeviceManagementResourceAccessProfileAssignmentIntentremove),
	}
}

func parseDeviceManagementResourceAccessProfileAssignmentIntent(input string) (*DeviceManagementResourceAccessProfileAssignmentIntent, error) {
	vals := map[string]DeviceManagementResourceAccessProfileAssignmentIntent{
		"apply":  DeviceManagementResourceAccessProfileAssignmentIntentapply,
		"remove": DeviceManagementResourceAccessProfileAssignmentIntentremove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementResourceAccessProfileAssignmentIntent(input)
	return &out, nil
}
