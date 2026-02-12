package checkprincipalaccess

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckPrincipalAccessRequest struct {
	Actions []RequiredAction `json:"actions"`
	Scope   string           `json:"scope"`
	Subject SubjectInfo      `json:"subject"`
}
