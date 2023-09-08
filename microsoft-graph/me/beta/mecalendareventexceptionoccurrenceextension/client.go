package mecalendareventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &MeCalendarEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
