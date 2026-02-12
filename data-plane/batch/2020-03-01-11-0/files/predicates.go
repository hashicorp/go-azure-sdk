package files

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeFileOperationPredicate struct {
	IsDirectory *bool
	Name        *string
	Url         *string
}

func (p NodeFileOperationPredicate) Matches(input NodeFile) bool {

	if p.IsDirectory != nil && (input.IsDirectory == nil || *p.IsDirectory != *input.IsDirectory) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
