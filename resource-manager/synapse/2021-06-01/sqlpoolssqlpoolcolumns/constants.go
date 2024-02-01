package sqlpoolssqlpoolcolumns

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ColumnDataType string

const (
	ColumnDataTypeBigint           ColumnDataType = "bigint"
	ColumnDataTypeBinary           ColumnDataType = "binary"
	ColumnDataTypeBit              ColumnDataType = "bit"
	ColumnDataTypeChar             ColumnDataType = "char"
	ColumnDataTypeDate             ColumnDataType = "date"
	ColumnDataTypeDatetime         ColumnDataType = "datetime"
	ColumnDataTypeDatetimeTwo      ColumnDataType = "datetime2"
	ColumnDataTypeDatetimeoffset   ColumnDataType = "datetimeoffset"
	ColumnDataTypeDecimal          ColumnDataType = "decimal"
	ColumnDataTypeFloat            ColumnDataType = "float"
	ColumnDataTypeGeography        ColumnDataType = "geography"
	ColumnDataTypeGeometry         ColumnDataType = "geometry"
	ColumnDataTypeHierarchyid      ColumnDataType = "hierarchyid"
	ColumnDataTypeImage            ColumnDataType = "image"
	ColumnDataTypeInt              ColumnDataType = "int"
	ColumnDataTypeMoney            ColumnDataType = "money"
	ColumnDataTypeNchar            ColumnDataType = "nchar"
	ColumnDataTypeNtext            ColumnDataType = "ntext"
	ColumnDataTypeNumeric          ColumnDataType = "numeric"
	ColumnDataTypeNvarchar         ColumnDataType = "nvarchar"
	ColumnDataTypeReal             ColumnDataType = "real"
	ColumnDataTypeSmalldatetime    ColumnDataType = "smalldatetime"
	ColumnDataTypeSmallint         ColumnDataType = "smallint"
	ColumnDataTypeSmallmoney       ColumnDataType = "smallmoney"
	ColumnDataTypeSqlVariant       ColumnDataType = "sql_variant"
	ColumnDataTypeSysname          ColumnDataType = "sysname"
	ColumnDataTypeText             ColumnDataType = "text"
	ColumnDataTypeTime             ColumnDataType = "time"
	ColumnDataTypeTimestamp        ColumnDataType = "timestamp"
	ColumnDataTypeTinyint          ColumnDataType = "tinyint"
	ColumnDataTypeUniqueidentifier ColumnDataType = "uniqueidentifier"
	ColumnDataTypeVarbinary        ColumnDataType = "varbinary"
	ColumnDataTypeVarchar          ColumnDataType = "varchar"
	ColumnDataTypeXml              ColumnDataType = "xml"
)

func PossibleValuesForColumnDataType() []string {
	return []string{
		string(ColumnDataTypeBigint),
		string(ColumnDataTypeBinary),
		string(ColumnDataTypeBit),
		string(ColumnDataTypeChar),
		string(ColumnDataTypeDate),
		string(ColumnDataTypeDatetime),
		string(ColumnDataTypeDatetimeTwo),
		string(ColumnDataTypeDatetimeoffset),
		string(ColumnDataTypeDecimal),
		string(ColumnDataTypeFloat),
		string(ColumnDataTypeGeography),
		string(ColumnDataTypeGeometry),
		string(ColumnDataTypeHierarchyid),
		string(ColumnDataTypeImage),
		string(ColumnDataTypeInt),
		string(ColumnDataTypeMoney),
		string(ColumnDataTypeNchar),
		string(ColumnDataTypeNtext),
		string(ColumnDataTypeNumeric),
		string(ColumnDataTypeNvarchar),
		string(ColumnDataTypeReal),
		string(ColumnDataTypeSmalldatetime),
		string(ColumnDataTypeSmallint),
		string(ColumnDataTypeSmallmoney),
		string(ColumnDataTypeSqlVariant),
		string(ColumnDataTypeSysname),
		string(ColumnDataTypeText),
		string(ColumnDataTypeTime),
		string(ColumnDataTypeTimestamp),
		string(ColumnDataTypeTinyint),
		string(ColumnDataTypeUniqueidentifier),
		string(ColumnDataTypeVarbinary),
		string(ColumnDataTypeVarchar),
		string(ColumnDataTypeXml),
	}
}

func (s *ColumnDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseColumnDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseColumnDataType(input string) (*ColumnDataType, error) {
	vals := map[string]ColumnDataType{
		"bigint":           ColumnDataTypeBigint,
		"binary":           ColumnDataTypeBinary,
		"bit":              ColumnDataTypeBit,
		"char":             ColumnDataTypeChar,
		"date":             ColumnDataTypeDate,
		"datetime":         ColumnDataTypeDatetime,
		"datetime2":        ColumnDataTypeDatetimeTwo,
		"datetimeoffset":   ColumnDataTypeDatetimeoffset,
		"decimal":          ColumnDataTypeDecimal,
		"float":            ColumnDataTypeFloat,
		"geography":        ColumnDataTypeGeography,
		"geometry":         ColumnDataTypeGeometry,
		"hierarchyid":      ColumnDataTypeHierarchyid,
		"image":            ColumnDataTypeImage,
		"int":              ColumnDataTypeInt,
		"money":            ColumnDataTypeMoney,
		"nchar":            ColumnDataTypeNchar,
		"ntext":            ColumnDataTypeNtext,
		"numeric":          ColumnDataTypeNumeric,
		"nvarchar":         ColumnDataTypeNvarchar,
		"real":             ColumnDataTypeReal,
		"smalldatetime":    ColumnDataTypeSmalldatetime,
		"smallint":         ColumnDataTypeSmallint,
		"smallmoney":       ColumnDataTypeSmallmoney,
		"sql_variant":      ColumnDataTypeSqlVariant,
		"sysname":          ColumnDataTypeSysname,
		"text":             ColumnDataTypeText,
		"time":             ColumnDataTypeTime,
		"timestamp":        ColumnDataTypeTimestamp,
		"tinyint":          ColumnDataTypeTinyint,
		"uniqueidentifier": ColumnDataTypeUniqueidentifier,
		"varbinary":        ColumnDataTypeVarbinary,
		"varchar":          ColumnDataTypeVarchar,
		"xml":              ColumnDataTypeXml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ColumnDataType(input)
	return &out, nil
}
