package integrationruntimeobjectmetadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SsisObjectMetadataOperationPredicate struct {
}

func (p SsisObjectMetadataOperationPredicate) Matches(input SsisObjectMetadata) bool {

	return true
}
