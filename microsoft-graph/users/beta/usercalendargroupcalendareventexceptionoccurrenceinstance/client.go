package usercalendargroupcalendareventexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
