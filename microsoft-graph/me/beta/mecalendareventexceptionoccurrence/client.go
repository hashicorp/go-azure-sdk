package mecalendareventexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExceptionOccurrenceClient: %+v", err)
	}

	return &MeCalendarEventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
