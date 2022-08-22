package diagnostic

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineDiagnosticSettings struct {
	Request  *HttpMessageDiagnostic `json:"request,omitempty"`
	Response *HttpMessageDiagnostic `json:"response,omitempty"`
}
