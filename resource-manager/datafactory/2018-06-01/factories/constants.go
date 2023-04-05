package factories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FactoryIdentityType string

const (
	FactoryIdentityTypeSystemAssigned             FactoryIdentityType = "SystemAssigned"
	FactoryIdentityTypeSystemAssignedUserAssigned FactoryIdentityType = "SystemAssigned,UserAssigned"
	FactoryIdentityTypeUserAssigned               FactoryIdentityType = "UserAssigned"
)

func PossibleValuesForFactoryIdentityType() []string {
	return []string{
		string(FactoryIdentityTypeSystemAssigned),
		string(FactoryIdentityTypeSystemAssignedUserAssigned),
		string(FactoryIdentityTypeUserAssigned),
	}
}

type GlobalParameterType string

const (
	GlobalParameterTypeArray  GlobalParameterType = "Array"
	GlobalParameterTypeBool   GlobalParameterType = "Bool"
	GlobalParameterTypeFloat  GlobalParameterType = "Float"
	GlobalParameterTypeInt    GlobalParameterType = "Int"
	GlobalParameterTypeObject GlobalParameterType = "Object"
	GlobalParameterTypeString GlobalParameterType = "String"
)

func PossibleValuesForGlobalParameterType() []string {
	return []string{
		string(GlobalParameterTypeArray),
		string(GlobalParameterTypeBool),
		string(GlobalParameterTypeFloat),
		string(GlobalParameterTypeInt),
		string(GlobalParameterTypeObject),
		string(GlobalParameterTypeString),
	}
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}
