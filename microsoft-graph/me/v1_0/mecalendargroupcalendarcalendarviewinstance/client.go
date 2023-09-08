package mecalendargroupcalendarcalendarviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewInstanceClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewInstanceClient{
		Client: client,
	}, nil
}
