package featurestoreentityversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndexColumn struct {
	ColumnName *string          `json:"columnName,omitempty"`
	DataType   *FeatureDataType `json:"dataType,omitempty"`
}
