package publishedartifact

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactOperationPredicate struct {
}

func (p ArtifactOperationPredicate) Matches(input Artifact) bool {

	return true
}
