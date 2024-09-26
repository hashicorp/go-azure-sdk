package siteonenotenotebooksectionpage

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

type CopySiteOnenoteNotebookSectionPageToSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenoteOperation
}

type CopySiteOnenoteNotebookSectionPageToSectionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCopySiteOnenoteNotebookSectionPageToSectionOperationOptions() CopySiteOnenoteNotebookSectionPageToSectionOperationOptions {
	return CopySiteOnenoteNotebookSectionPageToSectionOperationOptions{}
}

func (o CopySiteOnenoteNotebookSectionPageToSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopySiteOnenoteNotebookSectionPageToSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopySiteOnenoteNotebookSectionPageToSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopySiteOnenoteNotebookSectionPageToSection - Invoke action copyToSection. Copy a page to a specific section. For
// copy operations, you follow an asynchronous calling pattern: First call the Copy action, and then poll the operation
// endpoint for the result.
func (c SiteOnenoteNotebookSectionPageClient) CopySiteOnenoteNotebookSectionPageToSection(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookIdSectionIdPageId, input CopySiteOnenoteNotebookSectionPageToSectionRequest, options CopySiteOnenoteNotebookSectionPageToSectionOperationOptions) (result CopySiteOnenoteNotebookSectionPageToSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/copyToSection", id.ID()),
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

	var model stable.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
