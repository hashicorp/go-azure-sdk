package integrationserviceenvironmentnetworkhealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ErrorResponseCode string

const (
	ErrorResponseCodeIntegrationServiceEnvironmentNotFound ErrorResponseCode = "IntegrationServiceEnvironmentNotFound"
	ErrorResponseCodeInternalServerError                   ErrorResponseCode = "InternalServerError"
	ErrorResponseCodeInvalidOperationId                    ErrorResponseCode = "InvalidOperationId"
	ErrorResponseCodeNotSpecified                          ErrorResponseCode = "NotSpecified"
)

func PossibleValuesForErrorResponseCode() []string {
	return []string{
		string(ErrorResponseCodeIntegrationServiceEnvironmentNotFound),
		string(ErrorResponseCodeInternalServerError),
		string(ErrorResponseCodeInvalidOperationId),
		string(ErrorResponseCodeNotSpecified),
	}
}

type IntegrationServiceEnvironmentNetworkDependencyCategoryType string

const (
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAccessEndpoints                         IntegrationServiceEnvironmentNetworkDependencyCategoryType = "AccessEndpoints"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureActiveDirectory                    IntegrationServiceEnvironmentNetworkDependencyCategoryType = "AzureActiveDirectory"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureManagement                         IntegrationServiceEnvironmentNetworkDependencyCategoryType = "AzureManagement"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureStorage                            IntegrationServiceEnvironmentNetworkDependencyCategoryType = "AzureStorage"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeDiagnosticLogsAndMetrics                IntegrationServiceEnvironmentNetworkDependencyCategoryType = "DiagnosticLogsAndMetrics"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeIntegrationServiceEnvironmentConnectors IntegrationServiceEnvironmentNetworkDependencyCategoryType = "IntegrationServiceEnvironmentConnectors"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeNotSpecified                            IntegrationServiceEnvironmentNetworkDependencyCategoryType = "NotSpecified"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRecoveryService                         IntegrationServiceEnvironmentNetworkDependencyCategoryType = "RecoveryService"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRedisCache                              IntegrationServiceEnvironmentNetworkDependencyCategoryType = "RedisCache"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRegionalService                         IntegrationServiceEnvironmentNetworkDependencyCategoryType = "RegionalService"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSQL                                     IntegrationServiceEnvironmentNetworkDependencyCategoryType = "SQL"
	IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSSLCertificateVerification              IntegrationServiceEnvironmentNetworkDependencyCategoryType = "SSLCertificateVerification"
)

func PossibleValuesForIntegrationServiceEnvironmentNetworkDependencyCategoryType() []string {
	return []string{
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAccessEndpoints),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureActiveDirectory),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureManagement),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureStorage),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeDiagnosticLogsAndMetrics),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeIntegrationServiceEnvironmentConnectors),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeNotSpecified),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRecoveryService),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRedisCache),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRegionalService),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSQL),
		string(IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSSLCertificateVerification),
	}
}

type IntegrationServiceEnvironmentNetworkDependencyHealthState string

const (
	IntegrationServiceEnvironmentNetworkDependencyHealthStateHealthy      IntegrationServiceEnvironmentNetworkDependencyHealthState = "Healthy"
	IntegrationServiceEnvironmentNetworkDependencyHealthStateNotSpecified IntegrationServiceEnvironmentNetworkDependencyHealthState = "NotSpecified"
	IntegrationServiceEnvironmentNetworkDependencyHealthStateUnhealthy    IntegrationServiceEnvironmentNetworkDependencyHealthState = "Unhealthy"
	IntegrationServiceEnvironmentNetworkDependencyHealthStateUnknown      IntegrationServiceEnvironmentNetworkDependencyHealthState = "Unknown"
)

func PossibleValuesForIntegrationServiceEnvironmentNetworkDependencyHealthState() []string {
	return []string{
		string(IntegrationServiceEnvironmentNetworkDependencyHealthStateHealthy),
		string(IntegrationServiceEnvironmentNetworkDependencyHealthStateNotSpecified),
		string(IntegrationServiceEnvironmentNetworkDependencyHealthStateUnhealthy),
		string(IntegrationServiceEnvironmentNetworkDependencyHealthStateUnknown),
	}
}

type IntegrationServiceEnvironmentNetworkEndPointAccessibilityState string

const (
	IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable    IntegrationServiceEnvironmentNetworkEndPointAccessibilityState = "Available"
	IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateNotAvailable IntegrationServiceEnvironmentNetworkEndPointAccessibilityState = "NotAvailable"
	IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateNotSpecified IntegrationServiceEnvironmentNetworkEndPointAccessibilityState = "NotSpecified"
	IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateUnknown      IntegrationServiceEnvironmentNetworkEndPointAccessibilityState = "Unknown"
)

func PossibleValuesForIntegrationServiceEnvironmentNetworkEndPointAccessibilityState() []string {
	return []string{
		string(IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
		string(IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateNotAvailable),
		string(IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateNotSpecified),
		string(IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateUnknown),
	}
}
