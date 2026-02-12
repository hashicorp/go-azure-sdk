package pipelines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureFunctionActivityMethod string

const (
	AzureFunctionActivityMethodDELETE  AzureFunctionActivityMethod = "DELETE"
	AzureFunctionActivityMethodGET     AzureFunctionActivityMethod = "GET"
	AzureFunctionActivityMethodHEAD    AzureFunctionActivityMethod = "HEAD"
	AzureFunctionActivityMethodOPTIONS AzureFunctionActivityMethod = "OPTIONS"
	AzureFunctionActivityMethodPOST    AzureFunctionActivityMethod = "POST"
	AzureFunctionActivityMethodPUT     AzureFunctionActivityMethod = "PUT"
	AzureFunctionActivityMethodTRACE   AzureFunctionActivityMethod = "TRACE"
)

func PossibleValuesForAzureFunctionActivityMethod() []string {
	return []string{
		string(AzureFunctionActivityMethodDELETE),
		string(AzureFunctionActivityMethodGET),
		string(AzureFunctionActivityMethodHEAD),
		string(AzureFunctionActivityMethodOPTIONS),
		string(AzureFunctionActivityMethodPOST),
		string(AzureFunctionActivityMethodPUT),
		string(AzureFunctionActivityMethodTRACE),
	}
}

func (s *AzureFunctionActivityMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureFunctionActivityMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureFunctionActivityMethod(input string) (*AzureFunctionActivityMethod, error) {
	vals := map[string]AzureFunctionActivityMethod{
		"delete":  AzureFunctionActivityMethodDELETE,
		"get":     AzureFunctionActivityMethodGET,
		"head":    AzureFunctionActivityMethodHEAD,
		"options": AzureFunctionActivityMethodOPTIONS,
		"post":    AzureFunctionActivityMethodPOST,
		"put":     AzureFunctionActivityMethodPUT,
		"trace":   AzureFunctionActivityMethodTRACE,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureFunctionActivityMethod(input)
	return &out, nil
}

type AzureSearchIndexWriteBehaviorType string

const (
	AzureSearchIndexWriteBehaviorTypeMerge  AzureSearchIndexWriteBehaviorType = "Merge"
	AzureSearchIndexWriteBehaviorTypeUpload AzureSearchIndexWriteBehaviorType = "Upload"
)

func PossibleValuesForAzureSearchIndexWriteBehaviorType() []string {
	return []string{
		string(AzureSearchIndexWriteBehaviorTypeMerge),
		string(AzureSearchIndexWriteBehaviorTypeUpload),
	}
}

func (s *AzureSearchIndexWriteBehaviorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureSearchIndexWriteBehaviorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureSearchIndexWriteBehaviorType(input string) (*AzureSearchIndexWriteBehaviorType, error) {
	vals := map[string]AzureSearchIndexWriteBehaviorType{
		"merge":  AzureSearchIndexWriteBehaviorTypeMerge,
		"upload": AzureSearchIndexWriteBehaviorTypeUpload,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureSearchIndexWriteBehaviorType(input)
	return &out, nil
}

type CassandraSourceReadConsistencyLevels string

const (
	CassandraSourceReadConsistencyLevelsALL         CassandraSourceReadConsistencyLevels = "ALL"
	CassandraSourceReadConsistencyLevelsEACHQUORUM  CassandraSourceReadConsistencyLevels = "EACH_QUORUM"
	CassandraSourceReadConsistencyLevelsLOCALONE    CassandraSourceReadConsistencyLevels = "LOCAL_ONE"
	CassandraSourceReadConsistencyLevelsLOCALQUORUM CassandraSourceReadConsistencyLevels = "LOCAL_QUORUM"
	CassandraSourceReadConsistencyLevelsLOCALSERIAL CassandraSourceReadConsistencyLevels = "LOCAL_SERIAL"
	CassandraSourceReadConsistencyLevelsONE         CassandraSourceReadConsistencyLevels = "ONE"
	CassandraSourceReadConsistencyLevelsQUORUM      CassandraSourceReadConsistencyLevels = "QUORUM"
	CassandraSourceReadConsistencyLevelsSERIAL      CassandraSourceReadConsistencyLevels = "SERIAL"
	CassandraSourceReadConsistencyLevelsTHREE       CassandraSourceReadConsistencyLevels = "THREE"
	CassandraSourceReadConsistencyLevelsTWO         CassandraSourceReadConsistencyLevels = "TWO"
)

func PossibleValuesForCassandraSourceReadConsistencyLevels() []string {
	return []string{
		string(CassandraSourceReadConsistencyLevelsALL),
		string(CassandraSourceReadConsistencyLevelsEACHQUORUM),
		string(CassandraSourceReadConsistencyLevelsLOCALONE),
		string(CassandraSourceReadConsistencyLevelsLOCALQUORUM),
		string(CassandraSourceReadConsistencyLevelsLOCALSERIAL),
		string(CassandraSourceReadConsistencyLevelsONE),
		string(CassandraSourceReadConsistencyLevelsQUORUM),
		string(CassandraSourceReadConsistencyLevelsSERIAL),
		string(CassandraSourceReadConsistencyLevelsTHREE),
		string(CassandraSourceReadConsistencyLevelsTWO),
	}
}

func (s *CassandraSourceReadConsistencyLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCassandraSourceReadConsistencyLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCassandraSourceReadConsistencyLevels(input string) (*CassandraSourceReadConsistencyLevels, error) {
	vals := map[string]CassandraSourceReadConsistencyLevels{
		"all":          CassandraSourceReadConsistencyLevelsALL,
		"each_quorum":  CassandraSourceReadConsistencyLevelsEACHQUORUM,
		"local_one":    CassandraSourceReadConsistencyLevelsLOCALONE,
		"local_quorum": CassandraSourceReadConsistencyLevelsLOCALQUORUM,
		"local_serial": CassandraSourceReadConsistencyLevelsLOCALSERIAL,
		"one":          CassandraSourceReadConsistencyLevelsONE,
		"quorum":       CassandraSourceReadConsistencyLevelsQUORUM,
		"serial":       CassandraSourceReadConsistencyLevelsSERIAL,
		"three":        CassandraSourceReadConsistencyLevelsTHREE,
		"two":          CassandraSourceReadConsistencyLevelsTWO,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CassandraSourceReadConsistencyLevels(input)
	return &out, nil
}

type DataFlowComputeType string

const (
	DataFlowComputeTypeComputeOptimized DataFlowComputeType = "ComputeOptimized"
	DataFlowComputeTypeGeneral          DataFlowComputeType = "General"
	DataFlowComputeTypeMemoryOptimized  DataFlowComputeType = "MemoryOptimized"
)

func PossibleValuesForDataFlowComputeType() []string {
	return []string{
		string(DataFlowComputeTypeComputeOptimized),
		string(DataFlowComputeTypeGeneral),
		string(DataFlowComputeTypeMemoryOptimized),
	}
}

func (s *DataFlowComputeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataFlowComputeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataFlowComputeType(input string) (*DataFlowComputeType, error) {
	vals := map[string]DataFlowComputeType{
		"computeoptimized": DataFlowComputeTypeComputeOptimized,
		"general":          DataFlowComputeTypeGeneral,
		"memoryoptimized":  DataFlowComputeTypeMemoryOptimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataFlowComputeType(input)
	return &out, nil
}

type DataFlowReferenceType string

const (
	DataFlowReferenceTypeDataFlowReference DataFlowReferenceType = "DataFlowReference"
)

func PossibleValuesForDataFlowReferenceType() []string {
	return []string{
		string(DataFlowReferenceTypeDataFlowReference),
	}
}

func (s *DataFlowReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataFlowReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataFlowReferenceType(input string) (*DataFlowReferenceType, error) {
	vals := map[string]DataFlowReferenceType{
		"dataflowreference": DataFlowReferenceTypeDataFlowReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataFlowReferenceType(input)
	return &out, nil
}

type DatasetReferenceType string

const (
	DatasetReferenceTypeDatasetReference DatasetReferenceType = "DatasetReference"
)

func PossibleValuesForDatasetReferenceType() []string {
	return []string{
		string(DatasetReferenceTypeDatasetReference),
	}
}

func (s *DatasetReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatasetReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatasetReferenceType(input string) (*DatasetReferenceType, error) {
	vals := map[string]DatasetReferenceType{
		"datasetreference": DatasetReferenceTypeDatasetReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatasetReferenceType(input)
	return &out, nil
}

type DependencyCondition string

const (
	DependencyConditionCompleted DependencyCondition = "Completed"
	DependencyConditionFailed    DependencyCondition = "Failed"
	DependencyConditionSkipped   DependencyCondition = "Skipped"
	DependencyConditionSucceeded DependencyCondition = "Succeeded"
)

func PossibleValuesForDependencyCondition() []string {
	return []string{
		string(DependencyConditionCompleted),
		string(DependencyConditionFailed),
		string(DependencyConditionSkipped),
		string(DependencyConditionSucceeded),
	}
}

func (s *DependencyCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDependencyCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDependencyCondition(input string) (*DependencyCondition, error) {
	vals := map[string]DependencyCondition{
		"completed": DependencyConditionCompleted,
		"failed":    DependencyConditionFailed,
		"skipped":   DependencyConditionSkipped,
		"succeeded": DependencyConditionSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DependencyCondition(input)
	return &out, nil
}

type DynamicsSinkWriteBehavior string

const (
	DynamicsSinkWriteBehaviorUpsert DynamicsSinkWriteBehavior = "Upsert"
)

func PossibleValuesForDynamicsSinkWriteBehavior() []string {
	return []string{
		string(DynamicsSinkWriteBehaviorUpsert),
	}
}

func (s *DynamicsSinkWriteBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDynamicsSinkWriteBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDynamicsSinkWriteBehavior(input string) (*DynamicsSinkWriteBehavior, error) {
	vals := map[string]DynamicsSinkWriteBehavior{
		"upsert": DynamicsSinkWriteBehaviorUpsert,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DynamicsSinkWriteBehavior(input)
	return &out, nil
}

type ExpressionType string

const (
	ExpressionTypeExpression ExpressionType = "Expression"
)

func PossibleValuesForExpressionType() []string {
	return []string{
		string(ExpressionTypeExpression),
	}
}

func (s *ExpressionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExpressionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExpressionType(input string) (*ExpressionType, error) {
	vals := map[string]ExpressionType{
		"expression": ExpressionTypeExpression,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExpressionType(input)
	return &out, nil
}

type HDInsightActivityDebugInfoOption string

const (
	HDInsightActivityDebugInfoOptionAlways  HDInsightActivityDebugInfoOption = "Always"
	HDInsightActivityDebugInfoOptionFailure HDInsightActivityDebugInfoOption = "Failure"
	HDInsightActivityDebugInfoOptionNone    HDInsightActivityDebugInfoOption = "None"
)

func PossibleValuesForHDInsightActivityDebugInfoOption() []string {
	return []string{
		string(HDInsightActivityDebugInfoOptionAlways),
		string(HDInsightActivityDebugInfoOptionFailure),
		string(HDInsightActivityDebugInfoOptionNone),
	}
}

func (s *HDInsightActivityDebugInfoOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHDInsightActivityDebugInfoOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHDInsightActivityDebugInfoOption(input string) (*HDInsightActivityDebugInfoOption, error) {
	vals := map[string]HDInsightActivityDebugInfoOption{
		"always":  HDInsightActivityDebugInfoOptionAlways,
		"failure": HDInsightActivityDebugInfoOptionFailure,
		"none":    HDInsightActivityDebugInfoOptionNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HDInsightActivityDebugInfoOption(input)
	return &out, nil
}

type IntegrationRuntimeReferenceType string

const (
	IntegrationRuntimeReferenceTypeIntegrationRuntimeReference IntegrationRuntimeReferenceType = "IntegrationRuntimeReference"
)

func PossibleValuesForIntegrationRuntimeReferenceType() []string {
	return []string{
		string(IntegrationRuntimeReferenceTypeIntegrationRuntimeReference),
	}
}

func (s *IntegrationRuntimeReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntegrationRuntimeReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntegrationRuntimeReferenceType(input string) (*IntegrationRuntimeReferenceType, error) {
	vals := map[string]IntegrationRuntimeReferenceType{
		"integrationruntimereference": IntegrationRuntimeReferenceTypeIntegrationRuntimeReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationRuntimeReferenceType(input)
	return &out, nil
}

type JsonWriteFilePattern string

const (
	JsonWriteFilePatternArrayOfObjects JsonWriteFilePattern = "arrayOfObjects"
	JsonWriteFilePatternSetOfObjects   JsonWriteFilePattern = "setOfObjects"
)

func PossibleValuesForJsonWriteFilePattern() []string {
	return []string{
		string(JsonWriteFilePatternArrayOfObjects),
		string(JsonWriteFilePatternSetOfObjects),
	}
}

func (s *JsonWriteFilePattern) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJsonWriteFilePattern(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJsonWriteFilePattern(input string) (*JsonWriteFilePattern, error) {
	vals := map[string]JsonWriteFilePattern{
		"arrayofobjects": JsonWriteFilePatternArrayOfObjects,
		"setofobjects":   JsonWriteFilePatternSetOfObjects,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JsonWriteFilePattern(input)
	return &out, nil
}

type NetezzaPartitionOption string

const (
	NetezzaPartitionOptionDataSlice    NetezzaPartitionOption = "DataSlice"
	NetezzaPartitionOptionDynamicRange NetezzaPartitionOption = "DynamicRange"
	NetezzaPartitionOptionNone         NetezzaPartitionOption = "None"
)

func PossibleValuesForNetezzaPartitionOption() []string {
	return []string{
		string(NetezzaPartitionOptionDataSlice),
		string(NetezzaPartitionOptionDynamicRange),
		string(NetezzaPartitionOptionNone),
	}
}

func (s *NetezzaPartitionOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetezzaPartitionOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetezzaPartitionOption(input string) (*NetezzaPartitionOption, error) {
	vals := map[string]NetezzaPartitionOption{
		"dataslice":    NetezzaPartitionOptionDataSlice,
		"dynamicrange": NetezzaPartitionOptionDynamicRange,
		"none":         NetezzaPartitionOptionNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetezzaPartitionOption(input)
	return &out, nil
}

type NotebookReferenceType string

const (
	NotebookReferenceTypeNotebookReference NotebookReferenceType = "NotebookReference"
)

func PossibleValuesForNotebookReferenceType() []string {
	return []string{
		string(NotebookReferenceTypeNotebookReference),
	}
}

func (s *NotebookReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotebookReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotebookReferenceType(input string) (*NotebookReferenceType, error) {
	vals := map[string]NotebookReferenceType{
		"notebookreference": NotebookReferenceTypeNotebookReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotebookReferenceType(input)
	return &out, nil
}

type OraclePartitionOption string

const (
	OraclePartitionOptionDynamicRange              OraclePartitionOption = "DynamicRange"
	OraclePartitionOptionNone                      OraclePartitionOption = "None"
	OraclePartitionOptionPhysicalPartitionsOfTable OraclePartitionOption = "PhysicalPartitionsOfTable"
)

func PossibleValuesForOraclePartitionOption() []string {
	return []string{
		string(OraclePartitionOptionDynamicRange),
		string(OraclePartitionOptionNone),
		string(OraclePartitionOptionPhysicalPartitionsOfTable),
	}
}

func (s *OraclePartitionOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOraclePartitionOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOraclePartitionOption(input string) (*OraclePartitionOption, error) {
	vals := map[string]OraclePartitionOption{
		"dynamicrange":              OraclePartitionOptionDynamicRange,
		"none":                      OraclePartitionOptionNone,
		"physicalpartitionsoftable": OraclePartitionOptionPhysicalPartitionsOfTable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OraclePartitionOption(input)
	return &out, nil
}

type ParameterType string

const (
	ParameterTypeArray        ParameterType = "Array"
	ParameterTypeBool         ParameterType = "Bool"
	ParameterTypeFloat        ParameterType = "Float"
	ParameterTypeInt          ParameterType = "Int"
	ParameterTypeObject       ParameterType = "Object"
	ParameterTypeSecureString ParameterType = "SecureString"
	ParameterTypeString       ParameterType = "String"
)

func PossibleValuesForParameterType() []string {
	return []string{
		string(ParameterTypeArray),
		string(ParameterTypeBool),
		string(ParameterTypeFloat),
		string(ParameterTypeInt),
		string(ParameterTypeObject),
		string(ParameterTypeSecureString),
		string(ParameterTypeString),
	}
}

func (s *ParameterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseParameterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseParameterType(input string) (*ParameterType, error) {
	vals := map[string]ParameterType{
		"array":        ParameterTypeArray,
		"bool":         ParameterTypeBool,
		"float":        ParameterTypeFloat,
		"int":          ParameterTypeInt,
		"object":       ParameterTypeObject,
		"securestring": ParameterTypeSecureString,
		"string":       ParameterTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ParameterType(input)
	return &out, nil
}

type PipelineReferenceType string

const (
	PipelineReferenceTypePipelineReference PipelineReferenceType = "PipelineReference"
)

func PossibleValuesForPipelineReferenceType() []string {
	return []string{
		string(PipelineReferenceTypePipelineReference),
	}
}

func (s *PipelineReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePipelineReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePipelineReferenceType(input string) (*PipelineReferenceType, error) {
	vals := map[string]PipelineReferenceType{
		"pipelinereference": PipelineReferenceTypePipelineReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PipelineReferenceType(input)
	return &out, nil
}

type PolybaseSettingsRejectType string

const (
	PolybaseSettingsRejectTypePercentage PolybaseSettingsRejectType = "percentage"
	PolybaseSettingsRejectTypeValue      PolybaseSettingsRejectType = "value"
)

func PossibleValuesForPolybaseSettingsRejectType() []string {
	return []string{
		string(PolybaseSettingsRejectTypePercentage),
		string(PolybaseSettingsRejectTypeValue),
	}
}

func (s *PolybaseSettingsRejectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolybaseSettingsRejectType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolybaseSettingsRejectType(input string) (*PolybaseSettingsRejectType, error) {
	vals := map[string]PolybaseSettingsRejectType{
		"percentage": PolybaseSettingsRejectTypePercentage,
		"value":      PolybaseSettingsRejectTypeValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolybaseSettingsRejectType(input)
	return &out, nil
}

type SalesforceSinkWriteBehavior string

const (
	SalesforceSinkWriteBehaviorInsert SalesforceSinkWriteBehavior = "Insert"
	SalesforceSinkWriteBehaviorUpsert SalesforceSinkWriteBehavior = "Upsert"
)

func PossibleValuesForSalesforceSinkWriteBehavior() []string {
	return []string{
		string(SalesforceSinkWriteBehaviorInsert),
		string(SalesforceSinkWriteBehaviorUpsert),
	}
}

func (s *SalesforceSinkWriteBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSalesforceSinkWriteBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSalesforceSinkWriteBehavior(input string) (*SalesforceSinkWriteBehavior, error) {
	vals := map[string]SalesforceSinkWriteBehavior{
		"insert": SalesforceSinkWriteBehaviorInsert,
		"upsert": SalesforceSinkWriteBehaviorUpsert,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SalesforceSinkWriteBehavior(input)
	return &out, nil
}

type SalesforceSourceReadBehavior string

const (
	SalesforceSourceReadBehaviorQuery    SalesforceSourceReadBehavior = "Query"
	SalesforceSourceReadBehaviorQueryAll SalesforceSourceReadBehavior = "QueryAll"
)

func PossibleValuesForSalesforceSourceReadBehavior() []string {
	return []string{
		string(SalesforceSourceReadBehaviorQuery),
		string(SalesforceSourceReadBehaviorQueryAll),
	}
}

func (s *SalesforceSourceReadBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSalesforceSourceReadBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSalesforceSourceReadBehavior(input string) (*SalesforceSourceReadBehavior, error) {
	vals := map[string]SalesforceSourceReadBehavior{
		"query":    SalesforceSourceReadBehaviorQuery,
		"queryall": SalesforceSourceReadBehaviorQueryAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SalesforceSourceReadBehavior(input)
	return &out, nil
}

type SapCloudForCustomerSinkWriteBehavior string

const (
	SapCloudForCustomerSinkWriteBehaviorInsert SapCloudForCustomerSinkWriteBehavior = "Insert"
	SapCloudForCustomerSinkWriteBehaviorUpdate SapCloudForCustomerSinkWriteBehavior = "Update"
)

func PossibleValuesForSapCloudForCustomerSinkWriteBehavior() []string {
	return []string{
		string(SapCloudForCustomerSinkWriteBehaviorInsert),
		string(SapCloudForCustomerSinkWriteBehaviorUpdate),
	}
}

func (s *SapCloudForCustomerSinkWriteBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSapCloudForCustomerSinkWriteBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSapCloudForCustomerSinkWriteBehavior(input string) (*SapCloudForCustomerSinkWriteBehavior, error) {
	vals := map[string]SapCloudForCustomerSinkWriteBehavior{
		"insert": SapCloudForCustomerSinkWriteBehaviorInsert,
		"update": SapCloudForCustomerSinkWriteBehaviorUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SapCloudForCustomerSinkWriteBehavior(input)
	return &out, nil
}

type SapHanaPartitionOption string

const (
	SapHanaPartitionOptionNone                      SapHanaPartitionOption = "None"
	SapHanaPartitionOptionPhysicalPartitionsOfTable SapHanaPartitionOption = "PhysicalPartitionsOfTable"
	SapHanaPartitionOptionSapHanaDynamicRange       SapHanaPartitionOption = "SapHanaDynamicRange"
)

func PossibleValuesForSapHanaPartitionOption() []string {
	return []string{
		string(SapHanaPartitionOptionNone),
		string(SapHanaPartitionOptionPhysicalPartitionsOfTable),
		string(SapHanaPartitionOptionSapHanaDynamicRange),
	}
}

func (s *SapHanaPartitionOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSapHanaPartitionOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSapHanaPartitionOption(input string) (*SapHanaPartitionOption, error) {
	vals := map[string]SapHanaPartitionOption{
		"none":                      SapHanaPartitionOptionNone,
		"physicalpartitionsoftable": SapHanaPartitionOptionPhysicalPartitionsOfTable,
		"saphanadynamicrange":       SapHanaPartitionOptionSapHanaDynamicRange,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SapHanaPartitionOption(input)
	return &out, nil
}

type SapTablePartitionOption string

const (
	SapTablePartitionOptionNone                     SapTablePartitionOption = "None"
	SapTablePartitionOptionPartitionOnCalendarDate  SapTablePartitionOption = "PartitionOnCalendarDate"
	SapTablePartitionOptionPartitionOnCalendarMonth SapTablePartitionOption = "PartitionOnCalendarMonth"
	SapTablePartitionOptionPartitionOnCalendarYear  SapTablePartitionOption = "PartitionOnCalendarYear"
	SapTablePartitionOptionPartitionOnInt           SapTablePartitionOption = "PartitionOnInt"
	SapTablePartitionOptionPartitionOnTime          SapTablePartitionOption = "PartitionOnTime"
)

func PossibleValuesForSapTablePartitionOption() []string {
	return []string{
		string(SapTablePartitionOptionNone),
		string(SapTablePartitionOptionPartitionOnCalendarDate),
		string(SapTablePartitionOptionPartitionOnCalendarMonth),
		string(SapTablePartitionOptionPartitionOnCalendarYear),
		string(SapTablePartitionOptionPartitionOnInt),
		string(SapTablePartitionOptionPartitionOnTime),
	}
}

func (s *SapTablePartitionOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSapTablePartitionOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSapTablePartitionOption(input string) (*SapTablePartitionOption, error) {
	vals := map[string]SapTablePartitionOption{
		"none":                     SapTablePartitionOptionNone,
		"partitiononcalendardate":  SapTablePartitionOptionPartitionOnCalendarDate,
		"partitiononcalendarmonth": SapTablePartitionOptionPartitionOnCalendarMonth,
		"partitiononcalendaryear":  SapTablePartitionOptionPartitionOnCalendarYear,
		"partitiononint":           SapTablePartitionOptionPartitionOnInt,
		"partitionontime":          SapTablePartitionOptionPartitionOnTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SapTablePartitionOption(input)
	return &out, nil
}

type SparkJobReferenceType string

const (
	SparkJobReferenceTypeSparkJobDefinitionReference SparkJobReferenceType = "SparkJobDefinitionReference"
)

func PossibleValuesForSparkJobReferenceType() []string {
	return []string{
		string(SparkJobReferenceTypeSparkJobDefinitionReference),
	}
}

func (s *SparkJobReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSparkJobReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSparkJobReferenceType(input string) (*SparkJobReferenceType, error) {
	vals := map[string]SparkJobReferenceType{
		"sparkjobdefinitionreference": SparkJobReferenceTypeSparkJobDefinitionReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SparkJobReferenceType(input)
	return &out, nil
}

type SqlPoolReferenceType string

const (
	SqlPoolReferenceTypeSqlPoolReference SqlPoolReferenceType = "SqlPoolReference"
)

func PossibleValuesForSqlPoolReferenceType() []string {
	return []string{
		string(SqlPoolReferenceTypeSqlPoolReference),
	}
}

func (s *SqlPoolReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSqlPoolReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSqlPoolReferenceType(input string) (*SqlPoolReferenceType, error) {
	vals := map[string]SqlPoolReferenceType{
		"sqlpoolreference": SqlPoolReferenceTypeSqlPoolReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlPoolReferenceType(input)
	return &out, nil
}

type SsisLogLocationType string

const (
	SsisLogLocationTypeFile SsisLogLocationType = "File"
)

func PossibleValuesForSsisLogLocationType() []string {
	return []string{
		string(SsisLogLocationTypeFile),
	}
}

func (s *SsisLogLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSsisLogLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSsisLogLocationType(input string) (*SsisLogLocationType, error) {
	vals := map[string]SsisLogLocationType{
		"file": SsisLogLocationTypeFile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SsisLogLocationType(input)
	return &out, nil
}

type SsisPackageLocationType string

const (
	SsisPackageLocationTypeFile          SsisPackageLocationType = "File"
	SsisPackageLocationTypeInlinePackage SsisPackageLocationType = "InlinePackage"
	SsisPackageLocationTypePackageStore  SsisPackageLocationType = "PackageStore"
	SsisPackageLocationTypeSSISDB        SsisPackageLocationType = "SSISDB"
)

func PossibleValuesForSsisPackageLocationType() []string {
	return []string{
		string(SsisPackageLocationTypeFile),
		string(SsisPackageLocationTypeInlinePackage),
		string(SsisPackageLocationTypePackageStore),
		string(SsisPackageLocationTypeSSISDB),
	}
}

func (s *SsisPackageLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSsisPackageLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSsisPackageLocationType(input string) (*SsisPackageLocationType, error) {
	vals := map[string]SsisPackageLocationType{
		"file":          SsisPackageLocationTypeFile,
		"inlinepackage": SsisPackageLocationTypeInlinePackage,
		"packagestore":  SsisPackageLocationTypePackageStore,
		"ssisdb":        SsisPackageLocationTypeSSISDB,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SsisPackageLocationType(input)
	return &out, nil
}

type StoredProcedureParameterType string

const (
	StoredProcedureParameterTypeBoolean    StoredProcedureParameterType = "Boolean"
	StoredProcedureParameterTypeDate       StoredProcedureParameterType = "Date"
	StoredProcedureParameterTypeDecimal    StoredProcedureParameterType = "Decimal"
	StoredProcedureParameterTypeGuid       StoredProcedureParameterType = "Guid"
	StoredProcedureParameterTypeInt        StoredProcedureParameterType = "Int"
	StoredProcedureParameterTypeIntSixFour StoredProcedureParameterType = "Int64"
	StoredProcedureParameterTypeString     StoredProcedureParameterType = "String"
)

func PossibleValuesForStoredProcedureParameterType() []string {
	return []string{
		string(StoredProcedureParameterTypeBoolean),
		string(StoredProcedureParameterTypeDate),
		string(StoredProcedureParameterTypeDecimal),
		string(StoredProcedureParameterTypeGuid),
		string(StoredProcedureParameterTypeInt),
		string(StoredProcedureParameterTypeIntSixFour),
		string(StoredProcedureParameterTypeString),
	}
}

func (s *StoredProcedureParameterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStoredProcedureParameterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStoredProcedureParameterType(input string) (*StoredProcedureParameterType, error) {
	vals := map[string]StoredProcedureParameterType{
		"boolean": StoredProcedureParameterTypeBoolean,
		"date":    StoredProcedureParameterTypeDate,
		"decimal": StoredProcedureParameterTypeDecimal,
		"guid":    StoredProcedureParameterTypeGuid,
		"int":     StoredProcedureParameterTypeInt,
		"int64":   StoredProcedureParameterTypeIntSixFour,
		"string":  StoredProcedureParameterTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StoredProcedureParameterType(input)
	return &out, nil
}

type TeradataPartitionOption string

const (
	TeradataPartitionOptionDynamicRange TeradataPartitionOption = "DynamicRange"
	TeradataPartitionOptionHash         TeradataPartitionOption = "Hash"
	TeradataPartitionOptionNone         TeradataPartitionOption = "None"
)

func PossibleValuesForTeradataPartitionOption() []string {
	return []string{
		string(TeradataPartitionOptionDynamicRange),
		string(TeradataPartitionOptionHash),
		string(TeradataPartitionOptionNone),
	}
}

func (s *TeradataPartitionOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeradataPartitionOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeradataPartitionOption(input string) (*TeradataPartitionOption, error) {
	vals := map[string]TeradataPartitionOption{
		"dynamicrange": TeradataPartitionOptionDynamicRange,
		"hash":         TeradataPartitionOptionHash,
		"none":         TeradataPartitionOptionNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeradataPartitionOption(input)
	return &out, nil
}

type Type string

const (
	TypeLinkedServiceReference Type = "LinkedServiceReference"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeLinkedServiceReference),
	}
}

func (s *Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"linkedservicereference": TypeLinkedServiceReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}

type VariableType string

const (
	VariableTypeArray   VariableType = "Array"
	VariableTypeBool    VariableType = "Bool"
	VariableTypeBoolean VariableType = "Boolean"
	VariableTypeString  VariableType = "String"
)

func PossibleValuesForVariableType() []string {
	return []string{
		string(VariableTypeArray),
		string(VariableTypeBool),
		string(VariableTypeBoolean),
		string(VariableTypeString),
	}
}

func (s *VariableType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVariableType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVariableType(input string) (*VariableType, error) {
	vals := map[string]VariableType{
		"array":   VariableTypeArray,
		"bool":    VariableTypeBool,
		"boolean": VariableTypeBoolean,
		"string":  VariableTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VariableType(input)
	return &out, nil
}

type WebActivityMethod string

const (
	WebActivityMethodDELETE WebActivityMethod = "DELETE"
	WebActivityMethodGET    WebActivityMethod = "GET"
	WebActivityMethodPOST   WebActivityMethod = "POST"
	WebActivityMethodPUT    WebActivityMethod = "PUT"
)

func PossibleValuesForWebActivityMethod() []string {
	return []string{
		string(WebActivityMethodDELETE),
		string(WebActivityMethodGET),
		string(WebActivityMethodPOST),
		string(WebActivityMethodPUT),
	}
}

func (s *WebActivityMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebActivityMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebActivityMethod(input string) (*WebActivityMethod, error) {
	vals := map[string]WebActivityMethod{
		"delete": WebActivityMethodDELETE,
		"get":    WebActivityMethodGET,
		"post":   WebActivityMethodPOST,
		"put":    WebActivityMethodPUT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebActivityMethod(input)
	return &out, nil
}

type WebHookActivityMethod string

const (
	WebHookActivityMethodPOST WebHookActivityMethod = "POST"
)

func PossibleValuesForWebHookActivityMethod() []string {
	return []string{
		string(WebHookActivityMethodPOST),
	}
}

func (s *WebHookActivityMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebHookActivityMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebHookActivityMethod(input string) (*WebHookActivityMethod, error) {
	vals := map[string]WebHookActivityMethod{
		"post": WebHookActivityMethodPOST,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebHookActivityMethod(input)
	return &out, nil
}
