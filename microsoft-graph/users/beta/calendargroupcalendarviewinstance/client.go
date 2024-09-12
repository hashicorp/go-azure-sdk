package calendargroupcalendarviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewInstanceClient: %+v", err)
	}

	return &CalendarGroupCalendarViewInstanceClient{
		Client: client,
	}, nil
}
