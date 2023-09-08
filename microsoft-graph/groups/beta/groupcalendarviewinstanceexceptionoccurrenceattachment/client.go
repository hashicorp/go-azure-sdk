package groupcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
