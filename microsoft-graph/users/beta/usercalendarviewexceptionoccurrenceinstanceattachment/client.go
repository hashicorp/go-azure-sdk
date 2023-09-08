package usercalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &UserCalendarViewExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
