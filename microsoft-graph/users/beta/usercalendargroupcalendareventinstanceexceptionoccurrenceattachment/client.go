package usercalendargroupcalendareventinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
