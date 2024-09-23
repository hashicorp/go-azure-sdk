package calendarcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
