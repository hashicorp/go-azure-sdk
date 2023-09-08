package usercalendargroupcalendareventinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
