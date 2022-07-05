package streamingjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutputProperties struct {
	Datasource    *OutputDataSource `json:"datasource,omitempty"`
	Diagnostics   *Diagnostics      `json:"diagnostics,omitempty"`
	Etag          *string           `json:"etag,omitempty"`
	Serialization *Serialization    `json:"serialization,omitempty"`
	SizeWindow    *float64          `json:"sizeWindow,omitempty"`
	TimeWindow    *string           `json:"timeWindow,omitempty"`
}
