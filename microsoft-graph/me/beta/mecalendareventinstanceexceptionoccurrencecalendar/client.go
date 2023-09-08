package mecalendareventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeCalendarEventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
