package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedRunCommandScriptInput struct {
	ScriptParameters string `json:"scriptParameters"`
	ScriptURL        string `json:"scriptUrl"`
	StepName         string `json:"stepName"`
}
