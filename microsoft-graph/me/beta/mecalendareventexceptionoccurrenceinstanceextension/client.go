package mecalendareventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &MeCalendarEventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
