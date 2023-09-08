package mecalendareventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeCalendarEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
