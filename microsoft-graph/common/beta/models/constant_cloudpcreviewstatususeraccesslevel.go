package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCReviewStatusUserAccessLevel string

const (
	CloudPCReviewStatusUserAccessLevelrestricted   CloudPCReviewStatusUserAccessLevel = "Restricted"
	CloudPCReviewStatusUserAccessLevelunrestricted CloudPCReviewStatusUserAccessLevel = "Unrestricted"
)

func PossibleValuesForCloudPCReviewStatusUserAccessLevel() []string {
	return []string{
		string(CloudPCReviewStatusUserAccessLevelrestricted),
		string(CloudPCReviewStatusUserAccessLevelunrestricted),
	}
}

func parseCloudPCReviewStatusUserAccessLevel(input string) (*CloudPCReviewStatusUserAccessLevel, error) {
	vals := map[string]CloudPCReviewStatusUserAccessLevel{
		"restricted":   CloudPCReviewStatusUserAccessLevelrestricted,
		"unrestricted": CloudPCReviewStatusUserAccessLevelunrestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCReviewStatusUserAccessLevel(input)
	return &out, nil
}
