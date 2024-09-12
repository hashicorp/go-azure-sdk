package rejectedsender

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

type AddRejectedSenderRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// AddRejectedSenderRef - Create rejectedSender. Add a new user or group to the rejectedSender list. Specify the user or
// group in @odata.id in the request body. Users in the rejected senders list can't post to conversations of the group
// (identified in the POST request URL). Make sure you don't specify the same user or group in the rejected senders and
// accepted senders lists, otherwise you'll get an error.
func (c RejectedSenderClient) AddRejectedSenderRef(ctx context.Context, id beta.GroupId, input beta.ReferenceCreate) (result AddRejectedSenderRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/rejectedSenders/$ref", id.ID()),
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
