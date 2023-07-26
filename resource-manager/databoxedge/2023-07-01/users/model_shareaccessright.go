package users

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShareAccessRight struct {
	AccessType ShareAccessType `json:"accessType"`
	ShareId    string          `json:"shareId"`
}
