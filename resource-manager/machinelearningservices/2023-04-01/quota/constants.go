package quota

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaUnit string

const (
	QuotaUnitCount QuotaUnit = "Count"
)

func PossibleValuesForQuotaUnit() []string {
	return []string{
		string(QuotaUnitCount),
	}
}

func parseQuotaUnit(input string) (*QuotaUnit, error) {
	vals := map[string]QuotaUnit{
		"count": QuotaUnitCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QuotaUnit(input)
	return &out, nil
}

type Status string

const (
	StatusFailure                              Status = "Failure"
	StatusInvalidQuotaBelowClusterMinimum      Status = "InvalidQuotaBelowClusterMinimum"
	StatusInvalidQuotaExceedsSubscriptionLimit Status = "InvalidQuotaExceedsSubscriptionLimit"
	StatusInvalidVMFamilyName                  Status = "InvalidVMFamilyName"
	StatusOperationNotEnabledForRegion         Status = "OperationNotEnabledForRegion"
	StatusOperationNotSupportedForSku          Status = "OperationNotSupportedForSku"
	StatusSuccess                              Status = "Success"
	StatusUndefined                            Status = "Undefined"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusFailure),
		string(StatusInvalidQuotaBelowClusterMinimum),
		string(StatusInvalidQuotaExceedsSubscriptionLimit),
		string(StatusInvalidVMFamilyName),
		string(StatusOperationNotEnabledForRegion),
		string(StatusOperationNotSupportedForSku),
		string(StatusSuccess),
		string(StatusUndefined),
	}
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"failure":                              StatusFailure,
		"invalidquotabelowclusterminimum":      StatusInvalidQuotaBelowClusterMinimum,
		"invalidquotaexceedssubscriptionlimit": StatusInvalidQuotaExceedsSubscriptionLimit,
		"invalidvmfamilyname":                  StatusInvalidVMFamilyName,
		"operationnotenabledforregion":         StatusOperationNotEnabledForRegion,
		"operationnotsupportedforsku":          StatusOperationNotSupportedForSku,
		"success":                              StatusSuccess,
		"undefined":                            StatusUndefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
