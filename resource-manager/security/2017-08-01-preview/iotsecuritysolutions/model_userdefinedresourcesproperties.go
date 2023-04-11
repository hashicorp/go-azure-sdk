package iotsecuritysolutions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDefinedResourcesProperties struct {
	Query              string   `json:"query"`
	QuerySubscriptions []string `json:"querySubscriptions"`
}
