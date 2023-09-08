package usercalendareventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &UserCalendarEventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
