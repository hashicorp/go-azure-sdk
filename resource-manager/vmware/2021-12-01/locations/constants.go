package locations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaEnabled string

const (
	QuotaEnabledDisabled QuotaEnabled = "Disabled"
	QuotaEnabledEnabled  QuotaEnabled = "Enabled"
)

func PossibleValuesForQuotaEnabled() []string {
	return []string{
		string(QuotaEnabledDisabled),
		string(QuotaEnabledEnabled),
	}
}

type TrialStatus string

const (
	TrialStatusTrialAvailable TrialStatus = "TrialAvailable"
	TrialStatusTrialDisabled  TrialStatus = "TrialDisabled"
	TrialStatusTrialUsed      TrialStatus = "TrialUsed"
)

func PossibleValuesForTrialStatus() []string {
	return []string{
		string(TrialStatusTrialAvailable),
		string(TrialStatusTrialDisabled),
		string(TrialStatusTrialUsed),
	}
}
