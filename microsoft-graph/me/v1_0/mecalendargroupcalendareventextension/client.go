package mecalendargroupcalendareventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventExtensionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventExtensionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventExtensionClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventExtensionClient{
		Client: client,
	}, nil
}
