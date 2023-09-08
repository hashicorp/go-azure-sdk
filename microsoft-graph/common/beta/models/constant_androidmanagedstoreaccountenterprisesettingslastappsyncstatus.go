package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus string

const (
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusandroidForWorkApiError AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "AndroidForWorkApiError"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatuscredentialsNotValid    AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "CredentialsNotValid"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusmanagementServiceError AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "ManagementServiceError"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusnone                   AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "None"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatussuccess                AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "Success"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusunknownError           AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "UnknownError"
)

func PossibleValuesForAndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus() []string {
	return []string{
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusandroidForWorkApiError),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatuscredentialsNotValid),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusmanagementServiceError),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusnone),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatussuccess),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusunknownError),
	}
}

func parseAndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus(input string) (*AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus, error) {
	vals := map[string]AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus{
		"androidforworkapierror": AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusandroidForWorkApiError,
		"credentialsnotvalid":    AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatuscredentialsNotValid,
		"managementserviceerror": AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusmanagementServiceError,
		"none":                   AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusnone,
		"success":                AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatussuccess,
		"unknownerror":           AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatusunknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus(input)
	return &out, nil
}
