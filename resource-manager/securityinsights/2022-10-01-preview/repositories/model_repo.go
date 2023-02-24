package repositories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Repo struct {
	Branches *[]string `json:"branches,omitempty"`
	FullName *string   `json:"fullName,omitempty"`
	Url      *string   `json:"url,omitempty"`
}
