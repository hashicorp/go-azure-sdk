package mecalendargroupcalendareventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
