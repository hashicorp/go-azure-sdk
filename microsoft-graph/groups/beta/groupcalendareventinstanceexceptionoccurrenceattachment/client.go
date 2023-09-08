package groupcalendareventinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &GroupCalendarEventInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
