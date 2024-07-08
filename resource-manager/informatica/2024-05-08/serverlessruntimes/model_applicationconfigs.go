package serverlessruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationConfigs struct {
	Customized   string `json:"customized"`
	DefaultValue string `json:"defaultValue"`
	Name         string `json:"name"`
	Platform     string `json:"platform"`
	Type         string `json:"type"`
	Value        string `json:"value"`
}
