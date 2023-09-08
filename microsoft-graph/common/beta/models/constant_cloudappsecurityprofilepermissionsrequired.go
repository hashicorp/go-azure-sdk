package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecurityProfilePermissionsRequired string

const (
	CloudAppSecurityProfilePermissionsRequiredadministrator CloudAppSecurityProfilePermissionsRequired = "Administrator"
	CloudAppSecurityProfilePermissionsRequiredanonymous     CloudAppSecurityProfilePermissionsRequired = "Anonymous"
	CloudAppSecurityProfilePermissionsRequiredguest         CloudAppSecurityProfilePermissionsRequired = "Guest"
	CloudAppSecurityProfilePermissionsRequiredsystem        CloudAppSecurityProfilePermissionsRequired = "System"
	CloudAppSecurityProfilePermissionsRequiredunknown       CloudAppSecurityProfilePermissionsRequired = "Unknown"
	CloudAppSecurityProfilePermissionsRequireduser          CloudAppSecurityProfilePermissionsRequired = "User"
)

func PossibleValuesForCloudAppSecurityProfilePermissionsRequired() []string {
	return []string{
		string(CloudAppSecurityProfilePermissionsRequiredadministrator),
		string(CloudAppSecurityProfilePermissionsRequiredanonymous),
		string(CloudAppSecurityProfilePermissionsRequiredguest),
		string(CloudAppSecurityProfilePermissionsRequiredsystem),
		string(CloudAppSecurityProfilePermissionsRequiredunknown),
		string(CloudAppSecurityProfilePermissionsRequireduser),
	}
}

func parseCloudAppSecurityProfilePermissionsRequired(input string) (*CloudAppSecurityProfilePermissionsRequired, error) {
	vals := map[string]CloudAppSecurityProfilePermissionsRequired{
		"administrator": CloudAppSecurityProfilePermissionsRequiredadministrator,
		"anonymous":     CloudAppSecurityProfilePermissionsRequiredanonymous,
		"guest":         CloudAppSecurityProfilePermissionsRequiredguest,
		"system":        CloudAppSecurityProfilePermissionsRequiredsystem,
		"unknown":       CloudAppSecurityProfilePermissionsRequiredunknown,
		"user":          CloudAppSecurityProfilePermissionsRequireduser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudAppSecurityProfilePermissionsRequired(input)
	return &out, nil
}
