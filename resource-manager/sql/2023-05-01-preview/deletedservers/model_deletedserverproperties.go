package deletedservers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServerProperties struct {
	DeletionTime             *string `json:"deletionTime,omitempty"`
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`
	OriginalId               *string `json:"originalId,omitempty"`
	Version                  *string `json:"version,omitempty"`
}
