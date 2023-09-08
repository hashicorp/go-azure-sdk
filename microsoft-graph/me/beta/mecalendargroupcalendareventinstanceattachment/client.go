package mecalendargroupcalendareventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
