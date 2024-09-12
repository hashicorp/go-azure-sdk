package calendargroupcalendarevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventClient: %+v", err)
	}

	return &CalendarGroupCalendarEventClient{
		Client: client,
	}, nil
}
