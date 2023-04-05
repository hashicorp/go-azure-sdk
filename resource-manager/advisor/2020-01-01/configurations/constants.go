package configurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CPUThreshold string

const (
	CPUThresholdFive    CPUThreshold = "5"
	CPUThresholdOneFive CPUThreshold = "15"
	CPUThresholdOneZero CPUThreshold = "10"
	CPUThresholdTwoZero CPUThreshold = "20"
)

func PossibleValuesForCPUThreshold() []string {
	return []string{
		string(CPUThresholdFive),
		string(CPUThresholdOneFive),
		string(CPUThresholdOneZero),
		string(CPUThresholdTwoZero),
	}
}

type Category string

const (
	CategoryCost                  Category = "Cost"
	CategoryHighAvailability      Category = "HighAvailability"
	CategoryOperationalExcellence Category = "OperationalExcellence"
	CategoryPerformance           Category = "Performance"
	CategorySecurity              Category = "Security"
)

func PossibleValuesForCategory() []string {
	return []string{
		string(CategoryCost),
		string(CategoryHighAvailability),
		string(CategoryOperationalExcellence),
		string(CategoryPerformance),
		string(CategorySecurity),
	}
}

type DigestConfigState string

const (
	DigestConfigStateActive   DigestConfigState = "Active"
	DigestConfigStateDisabled DigestConfigState = "Disabled"
)

func PossibleValuesForDigestConfigState() []string {
	return []string{
		string(DigestConfigStateActive),
		string(DigestConfigStateDisabled),
	}
}
