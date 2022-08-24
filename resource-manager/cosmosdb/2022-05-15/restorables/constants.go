package restorables

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiType string

const (
	ApiTypeCassandra   ApiType = "Cassandra"
	ApiTypeGremlin     ApiType = "Gremlin"
	ApiTypeGremlinVTwo ApiType = "GremlinV2"
	ApiTypeMongoDB     ApiType = "MongoDB"
	ApiTypeSql         ApiType = "Sql"
	ApiTypeTable       ApiType = "Table"
)

func PossibleValuesForApiType() []string {
	return []string{
		string(ApiTypeCassandra),
		string(ApiTypeGremlin),
		string(ApiTypeGremlinVTwo),
		string(ApiTypeMongoDB),
		string(ApiTypeSql),
		string(ApiTypeTable),
	}
}

func parseApiType(input string) (*ApiType, error) {
	vals := map[string]ApiType{
		"cassandra": ApiTypeCassandra,
		"gremlin":   ApiTypeGremlin,
		"gremlinv2": ApiTypeGremlinVTwo,
		"mongodb":   ApiTypeMongoDB,
		"sql":       ApiTypeSql,
		"table":     ApiTypeTable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiType(input)
	return &out, nil
}

type OperationType string

const (
	OperationTypeCreate          OperationType = "Create"
	OperationTypeDelete          OperationType = "Delete"
	OperationTypeReplace         OperationType = "Replace"
	OperationTypeSystemOperation OperationType = "SystemOperation"
)

func PossibleValuesForOperationType() []string {
	return []string{
		string(OperationTypeCreate),
		string(OperationTypeDelete),
		string(OperationTypeReplace),
		string(OperationTypeSystemOperation),
	}
}

func parseOperationType(input string) (*OperationType, error) {
	vals := map[string]OperationType{
		"create":          OperationTypeCreate,
		"delete":          OperationTypeDelete,
		"replace":         OperationTypeReplace,
		"systemoperation": OperationTypeSystemOperation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationType(input)
	return &out, nil
}
