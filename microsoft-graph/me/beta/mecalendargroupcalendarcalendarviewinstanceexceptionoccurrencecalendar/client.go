package mecalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
