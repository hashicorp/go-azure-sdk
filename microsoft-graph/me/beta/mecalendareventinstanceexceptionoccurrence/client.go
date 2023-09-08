package mecalendareventinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &MeCalendarEventInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
