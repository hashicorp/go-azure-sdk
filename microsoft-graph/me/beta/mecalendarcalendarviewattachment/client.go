package mecalendarcalendarviewattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewAttachmentClient: %+v", err)
	}

	return &MeCalendarCalendarViewAttachmentClient{
		Client: client,
	}, nil
}
