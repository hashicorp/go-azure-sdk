package pendingexternaluserprofile

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

type DeletePendingExternalUserProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePendingExternalUserProfileOperationOptions struct {
	IfMatch *string
}

func DefaultDeletePendingExternalUserProfileOperationOptions() DeletePendingExternalUserProfileOperationOptions {
	return DeletePendingExternalUserProfileOperationOptions{}
}

func (o DeletePendingExternalUserProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePendingExternalUserProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeletePendingExternalUserProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePendingExternalUserProfile - Delete pendingExternalUserProfile. Delete a pendingExternalUserProfile object.
// Note: To permanently delete the pendingExternalUserProfile, follow permanently delete an item. To restore a
// pendingExternalUserProfile, follow restore a deleted item.
func (c PendingExternalUserProfileClient) DeletePendingExternalUserProfile(ctx context.Context, id beta.DirectoryPendingExternalUserProfileId, options DeletePendingExternalUserProfileOperationOptions) (result DeletePendingExternalUserProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
