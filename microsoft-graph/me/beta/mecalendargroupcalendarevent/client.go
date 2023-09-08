package mecalendargroupcalendarevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventClient{
		Client: client,
	}, nil
}
