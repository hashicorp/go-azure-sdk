package usercalendargroupcalendareventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
