package serverlessruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessRuntimeDependency struct {
	AppContextId    string `json:"appContextId"`
	Description     string `json:"description"`
	DocumentType    string `json:"documentType"`
	Id              string `json:"id"`
	LastUpdatedTime string `json:"lastUpdatedTime"`
	Path            string `json:"path"`
}
