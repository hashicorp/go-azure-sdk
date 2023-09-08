package mecalendargroupcalendareventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
