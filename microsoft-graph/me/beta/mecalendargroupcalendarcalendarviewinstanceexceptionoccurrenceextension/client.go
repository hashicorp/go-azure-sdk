package mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
