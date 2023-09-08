package mecalendargroupcalendareventexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventExceptionOccurrenceClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
