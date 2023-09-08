package mecalendargroupcalendarcalendarviewattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewAttachmentClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewAttachmentClient{
		Client: client,
	}, nil
}
