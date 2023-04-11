package externalsecuritysolutions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalSecuritySolutionOperationPredicate struct {
}

func (p ExternalSecuritySolutionOperationPredicate) Matches(input ExternalSecuritySolution) bool {

	return true
}
