package mecalendarcalendarviewinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarCalendarViewInstanceAttachmentClient{
		Client: client,
	}, nil
}
