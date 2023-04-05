package hypervmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HighlyAvailable string

const (
	HighlyAvailableNo      HighlyAvailable = "No"
	HighlyAvailableUnknown HighlyAvailable = "Unknown"
	HighlyAvailableYes     HighlyAvailable = "Yes"
)

func PossibleValuesForHighlyAvailable() []string {
	return []string{
		string(HighlyAvailableNo),
		string(HighlyAvailableUnknown),
		string(HighlyAvailableYes),
	}
}
