package calendareventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &CalendarEventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
