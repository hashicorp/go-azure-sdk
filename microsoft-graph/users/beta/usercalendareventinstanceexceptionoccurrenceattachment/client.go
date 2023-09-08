package usercalendareventinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &UserCalendarEventInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
