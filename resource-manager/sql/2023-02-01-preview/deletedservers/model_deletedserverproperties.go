package deletedservers

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServerProperties struct {
	DeletionTime             *string `json:"deletionTime,omitempty"`
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`
	OriginalId               *string `json:"originalId,omitempty"`
	Version                  *string `json:"version,omitempty"`
}

func (o *DeletedServerProperties) GetDeletionTimeAsTime() (*time.Time, error) {
	if o.DeletionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeletionTime, "2006-01-02T15:04:05Z07:00")
}
