package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceAttachmentPermission string

const (
	ReferenceAttachmentPermissionanonymousEdit    ReferenceAttachmentPermission = "AnonymousEdit"
	ReferenceAttachmentPermissionanonymousView    ReferenceAttachmentPermission = "AnonymousView"
	ReferenceAttachmentPermissionedit             ReferenceAttachmentPermission = "Edit"
	ReferenceAttachmentPermissionorganizationEdit ReferenceAttachmentPermission = "OrganizationEdit"
	ReferenceAttachmentPermissionorganizationView ReferenceAttachmentPermission = "OrganizationView"
	ReferenceAttachmentPermissionother            ReferenceAttachmentPermission = "Other"
	ReferenceAttachmentPermissionview             ReferenceAttachmentPermission = "View"
)

func PossibleValuesForReferenceAttachmentPermission() []string {
	return []string{
		string(ReferenceAttachmentPermissionanonymousEdit),
		string(ReferenceAttachmentPermissionanonymousView),
		string(ReferenceAttachmentPermissionedit),
		string(ReferenceAttachmentPermissionorganizationEdit),
		string(ReferenceAttachmentPermissionorganizationView),
		string(ReferenceAttachmentPermissionother),
		string(ReferenceAttachmentPermissionview),
	}
}

func parseReferenceAttachmentPermission(input string) (*ReferenceAttachmentPermission, error) {
	vals := map[string]ReferenceAttachmentPermission{
		"anonymousedit":    ReferenceAttachmentPermissionanonymousEdit,
		"anonymousview":    ReferenceAttachmentPermissionanonymousView,
		"edit":             ReferenceAttachmentPermissionedit,
		"organizationedit": ReferenceAttachmentPermissionorganizationEdit,
		"organizationview": ReferenceAttachmentPermissionorganizationView,
		"other":            ReferenceAttachmentPermissionother,
		"view":             ReferenceAttachmentPermissionview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReferenceAttachmentPermission(input)
	return &out, nil
}
