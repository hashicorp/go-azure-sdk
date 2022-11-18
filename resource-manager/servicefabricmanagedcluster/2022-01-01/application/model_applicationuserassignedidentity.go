package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationUserAssignedIdentity struct {
	Name        string `json:"name"`
	PrincipalId string `json:"principalId"`
}
