package dscnodeconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentSourceType string

const (
	ContentSourceTypeEmbeddedContent ContentSourceType = "embeddedContent"
	ContentSourceTypeUri             ContentSourceType = "uri"
)

func PossibleValuesForContentSourceType() []string {
	return []string{
		string(ContentSourceTypeEmbeddedContent),
		string(ContentSourceTypeUri),
	}
}
