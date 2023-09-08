package usercalendareventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &UserCalendarEventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
