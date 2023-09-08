package groupcalendarcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
