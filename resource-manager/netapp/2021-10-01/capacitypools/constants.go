package capacitypools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionType string

const (
	EncryptionTypeDouble EncryptionType = "Double"
	EncryptionTypeSingle EncryptionType = "Single"
)

func PossibleValuesForEncryptionType() []string {
	return []string{
		string(EncryptionTypeDouble),
		string(EncryptionTypeSingle),
	}
}

type QosType string

const (
	QosTypeAuto   QosType = "Auto"
	QosTypeManual QosType = "Manual"
)

func PossibleValuesForQosType() []string {
	return []string{
		string(QosTypeAuto),
		string(QosTypeManual),
	}
}

type ServiceLevel string

const (
	ServiceLevelPremium     ServiceLevel = "Premium"
	ServiceLevelStandard    ServiceLevel = "Standard"
	ServiceLevelStandardZRS ServiceLevel = "StandardZRS"
	ServiceLevelUltra       ServiceLevel = "Ultra"
)

func PossibleValuesForServiceLevel() []string {
	return []string{
		string(ServiceLevelPremium),
		string(ServiceLevelStandard),
		string(ServiceLevelStandardZRS),
		string(ServiceLevelUltra),
	}
}
