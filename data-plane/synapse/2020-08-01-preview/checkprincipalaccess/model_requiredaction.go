package checkprincipalaccess

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequiredAction struct {
	Id           string `json:"id"`
	IsDataAction bool   `json:"isDataAction"`
}
