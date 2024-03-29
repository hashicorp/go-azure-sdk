package forecast

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastProperties struct {
	Columns  *[]ForecastColumn `json:"columns,omitempty"`
	NextLink *string           `json:"nextLink,omitempty"`
	Rows     *[][]interface{}  `json:"rows,omitempty"`
}
