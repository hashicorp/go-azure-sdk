package calendareventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &CalendarEventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
