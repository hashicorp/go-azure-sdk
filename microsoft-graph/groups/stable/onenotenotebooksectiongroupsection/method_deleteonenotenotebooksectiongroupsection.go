package onenotenotebooksectiongroupsection

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

type DeleteOnenoteNotebookSectionGroupSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteOnenoteNotebookSectionGroupSectionOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteOnenoteNotebookSectionGroupSectionOperationOptions() DeleteOnenoteNotebookSectionGroupSectionOperationOptions {
	return DeleteOnenoteNotebookSectionGroupSectionOperationOptions{}
}

func (o DeleteOnenoteNotebookSectionGroupSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteOnenoteNotebookSectionGroupSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteOnenoteNotebookSectionGroupSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteOnenoteNotebookSectionGroupSection - Delete navigation property sections for groups
func (c OnenoteNotebookSectionGroupSectionClient) DeleteOnenoteNotebookSectionGroupSection(ctx context.Context, id stable.GroupIdOnenoteNotebookIdSectionGroupIdSectionId, options DeleteOnenoteNotebookSectionGroupSectionOperationOptions) (result DeleteOnenoteNotebookSectionGroupSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
