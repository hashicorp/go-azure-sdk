package attestation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuntimeData struct {
	Data     *string   `json:"data,omitempty"`
	DataType *DataType `json:"dataType,omitempty"`
}
