package usercalendarviewexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &UserCalendarViewExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
