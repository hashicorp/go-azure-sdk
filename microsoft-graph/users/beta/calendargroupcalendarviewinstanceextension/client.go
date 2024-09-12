package calendargroupcalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewInstanceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &CalendarGroupCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
