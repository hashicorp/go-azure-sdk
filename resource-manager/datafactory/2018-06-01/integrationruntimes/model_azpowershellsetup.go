package integrationruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzPowerShellSetup struct {
	Type           string                          `json:"type"`
	TypeProperties AzPowerShellSetupTypeProperties `json:"typeProperties"`
}
