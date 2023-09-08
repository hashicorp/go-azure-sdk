package groupcalendarviewexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &GroupCalendarViewExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
