package mecalendargroupcalendarcalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
