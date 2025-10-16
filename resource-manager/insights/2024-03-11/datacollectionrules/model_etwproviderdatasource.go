package datacollectionrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EtwProviderDataSource struct {
	EventIds     *[]string                           `json:"eventIds,omitempty"`
	Keyword      *string                             `json:"keyword,omitempty"`
	LogLevel     *KnownEtwProviderDataSourceLogLevel `json:"logLevel,omitempty"`
	Name         *string                             `json:"name,omitempty"`
	Provider     string                              `json:"provider"`
	ProviderType KnownEtwProviderType                `json:"providerType"`
	Streams      []string                            `json:"streams"`
}
