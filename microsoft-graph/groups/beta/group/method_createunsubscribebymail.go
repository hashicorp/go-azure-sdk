package group

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

type CreateUnsubscribeByMailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateUnsubscribeByMail - Invoke action unsubscribeByMail. Calling this method disables the current user to receive
// email notifications for this group about new posts, events, and files in that group. Supported for Microsoft 365
// groups only.
func (c GroupClient) CreateUnsubscribeByMail(ctx context.Context, id beta.GroupId) (result CreateUnsubscribeByMailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/unsubscribeByMail", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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
