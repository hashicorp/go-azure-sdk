package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WhatIfOperationProperties struct {
	Changes          *[]WhatIfChange                    `json:"changes,omitempty"`
	Diagnostics      *[]DeploymentDiagnosticsDefinition `json:"diagnostics,omitempty"`
	PotentialChanges *[]WhatIfChange                    `json:"potentialChanges,omitempty"`
}
