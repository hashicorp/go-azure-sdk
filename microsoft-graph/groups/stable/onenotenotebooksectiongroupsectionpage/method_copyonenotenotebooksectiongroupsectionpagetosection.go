package onenotenotebooksectiongroupsectionpage

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

type CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenoteOperation
}

type CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions() CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions {
	return CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions{}
}

func (o CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopyOnenoteNotebookSectionGroupSectionPageToSection - Invoke action copyToSection. Copy a page to a specific section.
// For copy operations, you follow an asynchronous calling pattern: First call the Copy action, and then poll the
// operation endpoint for the result.
func (c OnenoteNotebookSectionGroupSectionPageClient) CopyOnenoteNotebookSectionGroupSectionPageToSection(ctx context.Context, id stable.GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, input CopyOnenoteNotebookSectionGroupSectionPageToSectionRequest, options CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions) (result CopyOnenoteNotebookSectionGroupSectionPageToSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/copyToSection", id.ID()),
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

	var model stable.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
