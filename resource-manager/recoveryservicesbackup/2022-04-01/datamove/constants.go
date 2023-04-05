package datamove

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMoveLevel string

const (
	DataMoveLevelContainer DataMoveLevel = "Container"
	DataMoveLevelInvalid   DataMoveLevel = "Invalid"
	DataMoveLevelVault     DataMoveLevel = "Vault"
)

func PossibleValuesForDataMoveLevel() []string {
	return []string{
		string(DataMoveLevelContainer),
		string(DataMoveLevelInvalid),
		string(DataMoveLevelVault),
	}
}
