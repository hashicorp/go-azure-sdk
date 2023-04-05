package quota

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
