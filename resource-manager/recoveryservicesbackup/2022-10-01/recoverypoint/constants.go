package recoverypoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointTierType string

const (
	RecoveryPointTierTypeArchivedRP RecoveryPointTierType = "ArchivedRP"
	RecoveryPointTierTypeHardenedRP RecoveryPointTierType = "HardenedRP"
	RecoveryPointTierTypeInstantRP  RecoveryPointTierType = "InstantRP"
	RecoveryPointTierTypeInvalid    RecoveryPointTierType = "Invalid"
)

func PossibleValuesForRecoveryPointTierType() []string {
	return []string{
		string(RecoveryPointTierTypeArchivedRP),
		string(RecoveryPointTierTypeHardenedRP),
		string(RecoveryPointTierTypeInstantRP),
		string(RecoveryPointTierTypeInvalid),
	}
}
