package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessRuntimeConfigProperties struct {
	CdiConfigProps  *[]CdiConfigProps `json:"cdiConfigProps,omitempty"`
	CdieConfigProps *[]CdiConfigProps `json:"cdieConfigProps,omitempty"`
}
