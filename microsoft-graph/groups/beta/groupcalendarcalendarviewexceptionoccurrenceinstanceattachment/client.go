package groupcalendarcalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
