package sims

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimNameAndProperties struct {
	Name       string              `json:"name"`
	Properties SimPropertiesFormat `json:"properties"`
}
