package watchlists

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceType string

const (
	SourceTypeLocalFile     SourceType = "Local file"
	SourceTypeRemoteStorage SourceType = "Remote storage"
)

func PossibleValuesForSourceType() []string {
	return []string{
		string(SourceTypeLocalFile),
		string(SourceTypeRemoteStorage),
	}
}
