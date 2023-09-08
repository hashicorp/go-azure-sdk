package mecalendargroupcalendarcalendarviewextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewExtensionClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewExtensionClient{
		Client: client,
	}, nil
}
