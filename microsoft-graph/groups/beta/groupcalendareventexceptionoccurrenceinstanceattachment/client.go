package groupcalendareventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &GroupCalendarEventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
