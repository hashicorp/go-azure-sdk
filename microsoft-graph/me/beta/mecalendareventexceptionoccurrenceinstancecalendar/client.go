package mecalendareventexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &MeCalendarEventExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
