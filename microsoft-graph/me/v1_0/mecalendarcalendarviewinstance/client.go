package mecalendarcalendarviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewInstanceClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewInstanceClient: %+v", err)
	}

	return &MeCalendarCalendarViewInstanceClient{
		Client: client,
	}, nil
}
