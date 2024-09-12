package calendargroupcalendarviewattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewAttachmentClient: %+v", err)
	}

	return &CalendarGroupCalendarViewAttachmentClient{
		Client: client,
	}, nil
}
