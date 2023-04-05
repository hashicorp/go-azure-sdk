package integrationserviceenvironmentskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentSkuName string

const (
	IntegrationServiceEnvironmentSkuNameDeveloper    IntegrationServiceEnvironmentSkuName = "Developer"
	IntegrationServiceEnvironmentSkuNameNotSpecified IntegrationServiceEnvironmentSkuName = "NotSpecified"
	IntegrationServiceEnvironmentSkuNamePremium      IntegrationServiceEnvironmentSkuName = "Premium"
)

func PossibleValuesForIntegrationServiceEnvironmentSkuName() []string {
	return []string{
		string(IntegrationServiceEnvironmentSkuNameDeveloper),
		string(IntegrationServiceEnvironmentSkuNameNotSpecified),
		string(IntegrationServiceEnvironmentSkuNamePremium),
	}
}

type IntegrationServiceEnvironmentSkuScaleType string

const (
	IntegrationServiceEnvironmentSkuScaleTypeAutomatic IntegrationServiceEnvironmentSkuScaleType = "Automatic"
	IntegrationServiceEnvironmentSkuScaleTypeManual    IntegrationServiceEnvironmentSkuScaleType = "Manual"
	IntegrationServiceEnvironmentSkuScaleTypeNone      IntegrationServiceEnvironmentSkuScaleType = "None"
)

func PossibleValuesForIntegrationServiceEnvironmentSkuScaleType() []string {
	return []string{
		string(IntegrationServiceEnvironmentSkuScaleTypeAutomatic),
		string(IntegrationServiceEnvironmentSkuScaleTypeManual),
		string(IntegrationServiceEnvironmentSkuScaleTypeNone),
	}
}
