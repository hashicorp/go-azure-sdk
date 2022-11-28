package apidiagnostic

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineDiagnosticSettings struct {
	Request  *HTTPMessageDiagnostic `json:"request"`
	Response *HTTPMessageDiagnostic `json:"response"`
}
