package shares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccessRight struct {
	AccessType ShareAccessType `json:"accessType"`
	UserId     string          `json:"userId"`
}
