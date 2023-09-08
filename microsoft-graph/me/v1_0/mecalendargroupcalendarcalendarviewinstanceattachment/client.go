package mecalendargroupcalendarcalendarviewinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient{
		Client: client,
	}, nil
}
