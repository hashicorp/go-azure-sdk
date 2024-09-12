package deleteditem

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteDeletedItemOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDeletedItemOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteDeletedItemOperationOptions() DeleteDeletedItemOperationOptions {
	return DeleteDeletedItemOperationOptions{}
}

func (o DeleteDeletedItemOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDeletedItemOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteDeletedItemOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDeletedItem - Permanently delete an item (directory object). Permanently delete a recently deleted application,
// group, servicePrincipal, or user object from deleted items. After an item is permanently deleted, it cannot be
// restored. Administrative units cannot be permanently deleted by using the deletedItems API. Soft-deleted
// administrative units will be permanently deleted 30 days after initial deletion unless they are restored.
func (c DeletedItemClient) DeleteDeletedItem(ctx context.Context, id stable.DirectoryDeletedItemId, options DeleteDeletedItemOperationOptions) (result DeleteDeletedItemOperationResponse, err error) {
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
