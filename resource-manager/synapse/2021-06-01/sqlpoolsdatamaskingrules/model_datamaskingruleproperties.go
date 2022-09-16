package sqlpoolsdatamaskingrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingRuleProperties struct {
	AliasName         *string               `json:"aliasName,omitempty"`
	ColumnName        string                `json:"columnName"`
	Id                *string               `json:"id,omitempty"`
	MaskingFunction   DataMaskingFunction   `json:"maskingFunction"`
	NumberFrom        *string               `json:"numberFrom,omitempty"`
	NumberTo          *string               `json:"numberTo,omitempty"`
	PrefixSize        *string               `json:"prefixSize,omitempty"`
	ReplacementString *string               `json:"replacementString,omitempty"`
	RuleState         *DataMaskingRuleState `json:"ruleState,omitempty"`
	SchemaName        string                `json:"schemaName"`
	SuffixSize        *string               `json:"suffixSize,omitempty"`
	TableName         string                `json:"tableName"`
}
