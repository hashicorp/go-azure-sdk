package calendargroupcalendarcalendarviewinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarviewinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewInstanceAttachmentClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewInstanceAttachmentClient{
		Client: client,
	}, nil
}
