package mecalendargroupcalendareventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
