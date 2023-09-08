package mecalendargroupcalendareventinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventInstanceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventInstanceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventInstanceClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventInstanceClient{
		Client: client,
	}, nil
}
