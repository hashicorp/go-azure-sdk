package cosmosdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FullTextPolicy struct {
	DefaultLanguage *string         `json:"defaultLanguage,omitempty"`
	FullTextPaths   *[]FullTextPath `json:"fullTextPaths,omitempty"`
}
