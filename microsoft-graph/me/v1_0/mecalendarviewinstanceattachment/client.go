package mecalendarviewinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewInstanceAttachmentClient: %+v", err)
	}

	return &MeCalendarViewInstanceAttachmentClient{
		Client: client,
	}, nil
}
