package reservationsummaries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Datagrain string

const (
	DatagrainDaily   Datagrain = "daily"
	DatagrainMonthly Datagrain = "monthly"
)

func PossibleValuesForDatagrain() []string {
	return []string{
		string(DatagrainDaily),
		string(DatagrainMonthly),
	}
}
