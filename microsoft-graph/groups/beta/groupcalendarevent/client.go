package groupcalendarevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventClient: %+v", err)
	}

	return &GroupCalendarEventClient{
		Client: client,
	}, nil
}
