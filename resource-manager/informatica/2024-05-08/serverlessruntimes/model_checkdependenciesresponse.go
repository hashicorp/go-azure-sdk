package serverlessruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckDependenciesResponse struct {
	Count      int64                         `json:"count"`
	Id         string                        `json:"id"`
	References []ServerlessRuntimeDependency `json:"references"`
}
