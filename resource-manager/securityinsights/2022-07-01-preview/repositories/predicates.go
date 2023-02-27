package repositories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RepoOperationPredicate struct {
	FullName *string
	Url      *string
}

func (p RepoOperationPredicate) Matches(input Repo) bool {

	if p.FullName != nil && (input.FullName == nil && *p.FullName != *input.FullName) {
		return false
	}

	if p.Url != nil && (input.Url == nil && *p.Url != *input.Url) {
		return false
	}

	return true
}
