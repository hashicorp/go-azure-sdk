package integrationruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentVariableSetup struct {
	Type           string                                 `json:"type"`
	TypeProperties EnvironmentVariableSetupTypeProperties `json:"typeProperties"`
}
