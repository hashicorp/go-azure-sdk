package mecalendargroupcalendareventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
