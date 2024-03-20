package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureDatabricksDeltaLakeExportCommand struct {
	DateFormat      *interface{} `json:"dateFormat,omitempty"`
	TimestampFormat *interface{} `json:"timestampFormat,omitempty"`
	Type            string       `json:"type"`
}
