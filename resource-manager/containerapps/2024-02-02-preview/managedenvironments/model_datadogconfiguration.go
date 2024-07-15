package managedenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataDogConfiguration struct {
	Key  *string `json:"key,omitempty"`
	Site *string `json:"site,omitempty"`
}
