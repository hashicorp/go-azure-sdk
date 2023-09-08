package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecuritySessionControlCloudAppSecurityType string

const (
	CloudAppSecuritySessionControlCloudAppSecurityTypeblockDownloads CloudAppSecuritySessionControlCloudAppSecurityType = "BlockDownloads"
	CloudAppSecuritySessionControlCloudAppSecurityTypemcasConfigured CloudAppSecuritySessionControlCloudAppSecurityType = "McasConfigured"
	CloudAppSecuritySessionControlCloudAppSecurityTypemonitorOnly    CloudAppSecuritySessionControlCloudAppSecurityType = "MonitorOnly"
)

func PossibleValuesForCloudAppSecuritySessionControlCloudAppSecurityType() []string {
	return []string{
		string(CloudAppSecuritySessionControlCloudAppSecurityTypeblockDownloads),
		string(CloudAppSecuritySessionControlCloudAppSecurityTypemcasConfigured),
		string(CloudAppSecuritySessionControlCloudAppSecurityTypemonitorOnly),
	}
}

func parseCloudAppSecuritySessionControlCloudAppSecurityType(input string) (*CloudAppSecuritySessionControlCloudAppSecurityType, error) {
	vals := map[string]CloudAppSecuritySessionControlCloudAppSecurityType{
		"blockdownloads": CloudAppSecuritySessionControlCloudAppSecurityTypeblockDownloads,
		"mcasconfigured": CloudAppSecuritySessionControlCloudAppSecurityTypemcasConfigured,
		"monitoronly":    CloudAppSecuritySessionControlCloudAppSecurityTypemonitorOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudAppSecuritySessionControlCloudAppSecurityType(input)
	return &out, nil
}
