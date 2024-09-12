package acceptedsender

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddAcceptedSenderRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// AddAcceptedSenderRef - Create acceptedSender. Add a new user or group to the acceptedSender list. Specify the user or
// group in @odata.id in the request body. Users in the accepted senders list can post to conversations of the group.
// Make sure you don't specify the same user or group in the accepted senders and rejected senders lists, otherwise
// you'll get an error.
func (c AcceptedSenderClient) AddAcceptedSenderRef(ctx context.Context, id beta.GroupId, input beta.ReferenceCreate) (result AddAcceptedSenderRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/acceptedSenders/$ref", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
