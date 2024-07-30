package calendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &CalendarViewInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
