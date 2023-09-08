package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessEnrichedAuditLogsSettingsStatus string

const (
	NetworkaccessEnrichedAuditLogsSettingsStatusdisabled NetworkaccessEnrichedAuditLogsSettingsStatus = "Disabled"
	NetworkaccessEnrichedAuditLogsSettingsStatusenabled  NetworkaccessEnrichedAuditLogsSettingsStatus = "Enabled"
)

func PossibleValuesForNetworkaccessEnrichedAuditLogsSettingsStatus() []string {
	return []string{
		string(NetworkaccessEnrichedAuditLogsSettingsStatusdisabled),
		string(NetworkaccessEnrichedAuditLogsSettingsStatusenabled),
	}
}

func parseNetworkaccessEnrichedAuditLogsSettingsStatus(input string) (*NetworkaccessEnrichedAuditLogsSettingsStatus, error) {
	vals := map[string]NetworkaccessEnrichedAuditLogsSettingsStatus{
		"disabled": NetworkaccessEnrichedAuditLogsSettingsStatusdisabled,
		"enabled":  NetworkaccessEnrichedAuditLogsSettingsStatusenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessEnrichedAuditLogsSettingsStatus(input)
	return &out, nil
}
