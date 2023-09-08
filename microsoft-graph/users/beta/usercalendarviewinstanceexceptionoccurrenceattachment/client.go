package usercalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &UserCalendarViewInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
