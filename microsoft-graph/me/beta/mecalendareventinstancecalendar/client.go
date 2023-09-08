package mecalendareventinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventInstanceCalendarClient: %+v", err)
	}

	return &MeCalendarEventInstanceCalendarClient{
		Client: client,
	}, nil
}
