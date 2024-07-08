package serverlessruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CdiConfigProps struct {
	ApplicationConfigs []ApplicationConfigs `json:"applicationConfigs"`
	EngineName         string               `json:"engineName"`
	EngineVersion      string               `json:"engineVersion"`
}
