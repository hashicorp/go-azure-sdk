package mecalendargroupcalendareventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventAttachmentClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventAttachmentClient{
		Client: client,
	}, nil
}
