package mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
