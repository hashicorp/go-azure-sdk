package integrationserviceenvironmentmanagedapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiDeploymentParameterVisibility string

const (
	ApiDeploymentParameterVisibilityDefault      ApiDeploymentParameterVisibility = "Default"
	ApiDeploymentParameterVisibilityInternal     ApiDeploymentParameterVisibility = "Internal"
	ApiDeploymentParameterVisibilityNotSpecified ApiDeploymentParameterVisibility = "NotSpecified"
)

func PossibleValuesForApiDeploymentParameterVisibility() []string {
	return []string{
		string(ApiDeploymentParameterVisibilityDefault),
		string(ApiDeploymentParameterVisibilityInternal),
		string(ApiDeploymentParameterVisibilityNotSpecified),
	}
}

func (s *ApiDeploymentParameterVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApiDeploymentParameterVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApiDeploymentParameterVisibility(input string) (*ApiDeploymentParameterVisibility, error) {
	vals := map[string]ApiDeploymentParameterVisibility{
		"default":      ApiDeploymentParameterVisibilityDefault,
		"internal":     ApiDeploymentParameterVisibilityInternal,
		"notspecified": ApiDeploymentParameterVisibilityNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiDeploymentParameterVisibility(input)
	return &out, nil
}

type ApiTier string

const (
	ApiTierEnterprise   ApiTier = "Enterprise"
	ApiTierNotSpecified ApiTier = "NotSpecified"
	ApiTierPremium      ApiTier = "Premium"
	ApiTierStandard     ApiTier = "Standard"
)

func PossibleValuesForApiTier() []string {
	return []string{
		string(ApiTierEnterprise),
		string(ApiTierNotSpecified),
		string(ApiTierPremium),
		string(ApiTierStandard),
	}
}

func (s *ApiTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApiTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApiTier(input string) (*ApiTier, error) {
	vals := map[string]ApiTier{
		"enterprise":   ApiTierEnterprise,
		"notspecified": ApiTierNotSpecified,
		"premium":      ApiTierPremium,
		"standard":     ApiTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiTier(input)
	return &out, nil
}

type ApiType string

const (
	ApiTypeNotSpecified ApiType = "NotSpecified"
	ApiTypeRest         ApiType = "Rest"
	ApiTypeSoap         ApiType = "Soap"
)

func PossibleValuesForApiType() []string {
	return []string{
		string(ApiTypeNotSpecified),
		string(ApiTypeRest),
		string(ApiTypeSoap),
	}
}

func (s *ApiType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApiType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApiType(input string) (*ApiType, error) {
	vals := map[string]ApiType{
		"notspecified": ApiTypeNotSpecified,
		"rest":         ApiTypeRest,
		"soap":         ApiTypeSoap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiType(input)
	return &out, nil
}

type StatusAnnotation string

const (
	StatusAnnotationNotSpecified StatusAnnotation = "NotSpecified"
	StatusAnnotationPreview      StatusAnnotation = "Preview"
	StatusAnnotationProduction   StatusAnnotation = "Production"
)

func PossibleValuesForStatusAnnotation() []string {
	return []string{
		string(StatusAnnotationNotSpecified),
		string(StatusAnnotationPreview),
		string(StatusAnnotationProduction),
	}
}

func (s *StatusAnnotation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusAnnotation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusAnnotation(input string) (*StatusAnnotation, error) {
	vals := map[string]StatusAnnotation{
		"notspecified": StatusAnnotationNotSpecified,
		"preview":      StatusAnnotationPreview,
		"production":   StatusAnnotationProduction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusAnnotation(input)
	return &out, nil
}

type SwaggerSchemaType string

const (
	SwaggerSchemaTypeArray   SwaggerSchemaType = "Array"
	SwaggerSchemaTypeBoolean SwaggerSchemaType = "Boolean"
	SwaggerSchemaTypeFile    SwaggerSchemaType = "File"
	SwaggerSchemaTypeInteger SwaggerSchemaType = "Integer"
	SwaggerSchemaTypeNull    SwaggerSchemaType = "Null"
	SwaggerSchemaTypeNumber  SwaggerSchemaType = "Number"
	SwaggerSchemaTypeObject  SwaggerSchemaType = "Object"
	SwaggerSchemaTypeString  SwaggerSchemaType = "String"
)

func PossibleValuesForSwaggerSchemaType() []string {
	return []string{
		string(SwaggerSchemaTypeArray),
		string(SwaggerSchemaTypeBoolean),
		string(SwaggerSchemaTypeFile),
		string(SwaggerSchemaTypeInteger),
		string(SwaggerSchemaTypeNull),
		string(SwaggerSchemaTypeNumber),
		string(SwaggerSchemaTypeObject),
		string(SwaggerSchemaTypeString),
	}
}

func (s *SwaggerSchemaType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSwaggerSchemaType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSwaggerSchemaType(input string) (*SwaggerSchemaType, error) {
	vals := map[string]SwaggerSchemaType{
		"array":   SwaggerSchemaTypeArray,
		"boolean": SwaggerSchemaTypeBoolean,
		"file":    SwaggerSchemaTypeFile,
		"integer": SwaggerSchemaTypeInteger,
		"null":    SwaggerSchemaTypeNull,
		"number":  SwaggerSchemaTypeNumber,
		"object":  SwaggerSchemaTypeObject,
		"string":  SwaggerSchemaTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SwaggerSchemaType(input)
	return &out, nil
}

type WorkflowProvisioningState string

const (
	WorkflowProvisioningStateAccepted      WorkflowProvisioningState = "Accepted"
	WorkflowProvisioningStateCanceled      WorkflowProvisioningState = "Canceled"
	WorkflowProvisioningStateCompleted     WorkflowProvisioningState = "Completed"
	WorkflowProvisioningStateCreated       WorkflowProvisioningState = "Created"
	WorkflowProvisioningStateCreating      WorkflowProvisioningState = "Creating"
	WorkflowProvisioningStateDeleted       WorkflowProvisioningState = "Deleted"
	WorkflowProvisioningStateDeleting      WorkflowProvisioningState = "Deleting"
	WorkflowProvisioningStateFailed        WorkflowProvisioningState = "Failed"
	WorkflowProvisioningStateInProgress    WorkflowProvisioningState = "InProgress"
	WorkflowProvisioningStateMoving        WorkflowProvisioningState = "Moving"
	WorkflowProvisioningStateNotSpecified  WorkflowProvisioningState = "NotSpecified"
	WorkflowProvisioningStatePending       WorkflowProvisioningState = "Pending"
	WorkflowProvisioningStateReady         WorkflowProvisioningState = "Ready"
	WorkflowProvisioningStateRegistered    WorkflowProvisioningState = "Registered"
	WorkflowProvisioningStateRegistering   WorkflowProvisioningState = "Registering"
	WorkflowProvisioningStateRenewing      WorkflowProvisioningState = "Renewing"
	WorkflowProvisioningStateRunning       WorkflowProvisioningState = "Running"
	WorkflowProvisioningStateSucceeded     WorkflowProvisioningState = "Succeeded"
	WorkflowProvisioningStateUnregistered  WorkflowProvisioningState = "Unregistered"
	WorkflowProvisioningStateUnregistering WorkflowProvisioningState = "Unregistering"
	WorkflowProvisioningStateUpdating      WorkflowProvisioningState = "Updating"
	WorkflowProvisioningStateWaiting       WorkflowProvisioningState = "Waiting"
)

func PossibleValuesForWorkflowProvisioningState() []string {
	return []string{
		string(WorkflowProvisioningStateAccepted),
		string(WorkflowProvisioningStateCanceled),
		string(WorkflowProvisioningStateCompleted),
		string(WorkflowProvisioningStateCreated),
		string(WorkflowProvisioningStateCreating),
		string(WorkflowProvisioningStateDeleted),
		string(WorkflowProvisioningStateDeleting),
		string(WorkflowProvisioningStateFailed),
		string(WorkflowProvisioningStateInProgress),
		string(WorkflowProvisioningStateMoving),
		string(WorkflowProvisioningStateNotSpecified),
		string(WorkflowProvisioningStatePending),
		string(WorkflowProvisioningStateReady),
		string(WorkflowProvisioningStateRegistered),
		string(WorkflowProvisioningStateRegistering),
		string(WorkflowProvisioningStateRenewing),
		string(WorkflowProvisioningStateRunning),
		string(WorkflowProvisioningStateSucceeded),
		string(WorkflowProvisioningStateUnregistered),
		string(WorkflowProvisioningStateUnregistering),
		string(WorkflowProvisioningStateUpdating),
		string(WorkflowProvisioningStateWaiting),
	}
}

func (s *WorkflowProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkflowProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkflowProvisioningState(input string) (*WorkflowProvisioningState, error) {
	vals := map[string]WorkflowProvisioningState{
		"accepted":      WorkflowProvisioningStateAccepted,
		"canceled":      WorkflowProvisioningStateCanceled,
		"completed":     WorkflowProvisioningStateCompleted,
		"created":       WorkflowProvisioningStateCreated,
		"creating":      WorkflowProvisioningStateCreating,
		"deleted":       WorkflowProvisioningStateDeleted,
		"deleting":      WorkflowProvisioningStateDeleting,
		"failed":        WorkflowProvisioningStateFailed,
		"inprogress":    WorkflowProvisioningStateInProgress,
		"moving":        WorkflowProvisioningStateMoving,
		"notspecified":  WorkflowProvisioningStateNotSpecified,
		"pending":       WorkflowProvisioningStatePending,
		"ready":         WorkflowProvisioningStateReady,
		"registered":    WorkflowProvisioningStateRegistered,
		"registering":   WorkflowProvisioningStateRegistering,
		"renewing":      WorkflowProvisioningStateRenewing,
		"running":       WorkflowProvisioningStateRunning,
		"succeeded":     WorkflowProvisioningStateSucceeded,
		"unregistered":  WorkflowProvisioningStateUnregistered,
		"unregistering": WorkflowProvisioningStateUnregistering,
		"updating":      WorkflowProvisioningStateUpdating,
		"waiting":       WorkflowProvisioningStateWaiting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowProvisioningState(input)
	return &out, nil
}

type WsdlImportMethod string

const (
	WsdlImportMethodNotSpecified    WsdlImportMethod = "NotSpecified"
	WsdlImportMethodSoapPassThrough WsdlImportMethod = "SoapPassThrough"
	WsdlImportMethodSoapToRest      WsdlImportMethod = "SoapToRest"
)

func PossibleValuesForWsdlImportMethod() []string {
	return []string{
		string(WsdlImportMethodNotSpecified),
		string(WsdlImportMethodSoapPassThrough),
		string(WsdlImportMethodSoapToRest),
	}
}

func (s *WsdlImportMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWsdlImportMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWsdlImportMethod(input string) (*WsdlImportMethod, error) {
	vals := map[string]WsdlImportMethod{
		"notspecified":    WsdlImportMethodNotSpecified,
		"soappassthrough": WsdlImportMethodSoapPassThrough,
		"soaptorest":      WsdlImportMethodSoapToRest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WsdlImportMethod(input)
	return &out, nil
}
