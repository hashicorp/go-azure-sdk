package onenotenotebooksectiongroupsection

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

type CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.OnenoteOperation
}

type CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions() CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions {
	return CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions{}
}

func (o CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopyOnenoteNotebookSectionGroupSectionToSectionGroup - Invoke action copyToSectionGroup. Copies a section to a
// specific section group. For Copy operations, you follow an asynchronous calling pattern: First call the Copy action,
// and then poll the operation endpoint for the result.
func (c OnenoteNotebookSectionGroupSectionClient) CopyOnenoteNotebookSectionGroupSectionToSectionGroup(ctx context.Context, id beta.MeOnenoteNotebookIdSectionGroupIdSectionId, input CopyOnenoteNotebookSectionGroupSectionToSectionGroupRequest, options CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions) (result CopyOnenoteNotebookSectionGroupSectionToSectionGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/copyToSectionGroup", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
