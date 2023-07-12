package integrationserviceenvironmentnetworkhealth

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *ErrorResponseCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseErrorResponseCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseErrorResponseCode(input string) (*ErrorResponseCode, error) {
	vals := map[string]ErrorResponseCode{
		"integrationserviceenvironmentnotfound": ErrorResponseCodeIntegrationServiceEnvironmentNotFound,
		"internalservererror":                   ErrorResponseCodeInternalServerError,
		"invalidoperationid":                    ErrorResponseCodeInvalidOperationId,
		"notspecified":                          ErrorResponseCodeNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ErrorResponseCode(input)
	return &out, nil
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

func (s *IntegrationServiceEnvironmentNetworkDependencyCategoryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntegrationServiceEnvironmentNetworkDependencyCategoryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntegrationServiceEnvironmentNetworkDependencyCategoryType(input string) (*IntegrationServiceEnvironmentNetworkDependencyCategoryType, error) {
	vals := map[string]IntegrationServiceEnvironmentNetworkDependencyCategoryType{
		"accessendpoints":                         IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAccessEndpoints,
		"azureactivedirectory":                    IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureActiveDirectory,
		"azuremanagement":                         IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureManagement,
		"azurestorage":                            IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureStorage,
		"diagnosticlogsandmetrics":                IntegrationServiceEnvironmentNetworkDependencyCategoryTypeDiagnosticLogsAndMetrics,
		"integrationserviceenvironmentconnectors": IntegrationServiceEnvironmentNetworkDependencyCategoryTypeIntegrationServiceEnvironmentConnectors,
		"notspecified":                            IntegrationServiceEnvironmentNetworkDependencyCategoryTypeNotSpecified,
		"recoveryservice":                         IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRecoveryService,
		"rediscache":                              IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRedisCache,
		"regionalservice":                         IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRegionalService,
		"sql":                                     IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSQL,
		"sslcertificateverification":              IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSSLCertificateVerification,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationServiceEnvironmentNetworkDependencyCategoryType(input)
	return &out, nil
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

func (s *IntegrationServiceEnvironmentNetworkDependencyHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntegrationServiceEnvironmentNetworkDependencyHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntegrationServiceEnvironmentNetworkDependencyHealthState(input string) (*IntegrationServiceEnvironmentNetworkDependencyHealthState, error) {
	vals := map[string]IntegrationServiceEnvironmentNetworkDependencyHealthState{
		"healthy":      IntegrationServiceEnvironmentNetworkDependencyHealthStateHealthy,
		"notspecified": IntegrationServiceEnvironmentNetworkDependencyHealthStateNotSpecified,
		"unhealthy":    IntegrationServiceEnvironmentNetworkDependencyHealthStateUnhealthy,
		"unknown":      IntegrationServiceEnvironmentNetworkDependencyHealthStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationServiceEnvironmentNetworkDependencyHealthState(input)
	return &out, nil
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

func (s *IntegrationServiceEnvironmentNetworkEndPointAccessibilityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntegrationServiceEnvironmentNetworkEndPointAccessibilityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntegrationServiceEnvironmentNetworkEndPointAccessibilityState(input string) (*IntegrationServiceEnvironmentNetworkEndPointAccessibilityState, error) {
	vals := map[string]IntegrationServiceEnvironmentNetworkEndPointAccessibilityState{
		"available":    IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable,
		"notavailable": IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateNotAvailable,
		"notspecified": IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateNotSpecified,
		"unknown":      IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationServiceEnvironmentNetworkEndPointAccessibilityState(input)
	return &out, nil
}
