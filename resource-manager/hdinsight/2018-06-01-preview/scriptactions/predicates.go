package scriptactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuntimeScriptActionOperationPredicate struct {
	ApplicationName *string
	Name            *string
	Parameters      *string
	Uri             *string
}

func (p RuntimeScriptActionOperationPredicate) Matches(input RuntimeScriptAction) bool {

	if p.ApplicationName != nil && (input.ApplicationName == nil || *p.ApplicationName != *input.ApplicationName) {
		return false
	}

	if p.Name != nil && *p.Name != input.Name {
		return false
	}

	if p.Parameters != nil && (input.Parameters == nil || *p.Parameters != *input.Parameters) {
		return false
	}

	if p.Uri != nil && *p.Uri != input.Uri {
		return false
	}

	return true
}
