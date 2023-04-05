package nodecountinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CountType string

const (
	CountTypeNodeconfiguration CountType = "nodeconfiguration"
	CountTypeStatus            CountType = "status"
)

func PossibleValuesForCountType() []string {
	return []string{
		string(CountTypeNodeconfiguration),
		string(CountTypeStatus),
	}
}
