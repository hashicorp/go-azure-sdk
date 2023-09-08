package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalculatedColumn struct {
	Format     *string `json:"format,omitempty"`
	Formula    *string `json:"formula,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	OutputType *string `json:"outputType,omitempty"`
}
