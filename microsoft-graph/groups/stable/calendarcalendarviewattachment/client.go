package calendarcalendarviewattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewAttachmentClient: %+v", err)
	}

	return &CalendarCalendarViewAttachmentClient{
		Client: client,
	}, nil
}
