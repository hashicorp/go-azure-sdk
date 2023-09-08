package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSettingsLastAppSyncStatus string

const (
	AndroidForWorkSettingsLastAppSyncStatusandroidForWorkApiError AndroidForWorkSettingsLastAppSyncStatus = "AndroidForWorkApiError"
	AndroidForWorkSettingsLastAppSyncStatuscredentialsNotValid    AndroidForWorkSettingsLastAppSyncStatus = "CredentialsNotValid"
	AndroidForWorkSettingsLastAppSyncStatusmanagementServiceError AndroidForWorkSettingsLastAppSyncStatus = "ManagementServiceError"
	AndroidForWorkSettingsLastAppSyncStatusnone                   AndroidForWorkSettingsLastAppSyncStatus = "None"
	AndroidForWorkSettingsLastAppSyncStatussuccess                AndroidForWorkSettingsLastAppSyncStatus = "Success"
	AndroidForWorkSettingsLastAppSyncStatusunknownError           AndroidForWorkSettingsLastAppSyncStatus = "UnknownError"
)

func PossibleValuesForAndroidForWorkSettingsLastAppSyncStatus() []string {
	return []string{
		string(AndroidForWorkSettingsLastAppSyncStatusandroidForWorkApiError),
		string(AndroidForWorkSettingsLastAppSyncStatuscredentialsNotValid),
		string(AndroidForWorkSettingsLastAppSyncStatusmanagementServiceError),
		string(AndroidForWorkSettingsLastAppSyncStatusnone),
		string(AndroidForWorkSettingsLastAppSyncStatussuccess),
		string(AndroidForWorkSettingsLastAppSyncStatusunknownError),
	}
}

func parseAndroidForWorkSettingsLastAppSyncStatus(input string) (*AndroidForWorkSettingsLastAppSyncStatus, error) {
	vals := map[string]AndroidForWorkSettingsLastAppSyncStatus{
		"androidforworkapierror": AndroidForWorkSettingsLastAppSyncStatusandroidForWorkApiError,
		"credentialsnotvalid":    AndroidForWorkSettingsLastAppSyncStatuscredentialsNotValid,
		"managementserviceerror": AndroidForWorkSettingsLastAppSyncStatusmanagementServiceError,
		"none":                   AndroidForWorkSettingsLastAppSyncStatusnone,
		"success":                AndroidForWorkSettingsLastAppSyncStatussuccess,
		"unknownerror":           AndroidForWorkSettingsLastAppSyncStatusunknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkSettingsLastAppSyncStatus(input)
	return &out, nil
}
