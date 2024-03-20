package datasets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomDataset struct {
	Annotations       *[]interface{}                     `json:"annotations,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Folder            *DatasetFolder                     `json:"folder,omitempty"`
	LinkedServiceName LinkedServiceReference             `json:"linkedServiceName"`
	Parameters        *map[string]ParameterSpecification `json:"parameters,omitempty"`
	Schema            *interface{}                       `json:"schema,omitempty"`
	Structure         *interface{}                       `json:"structure,omitempty"`
	Type              string                             `json:"type"`
	TypeProperties    *interface{}                       `json:"typeProperties,omitempty"`
}
