package siteonenotenotebooksection

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

type CreateSiteOnenoteNotebookSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenoteSection
}

type CreateSiteOnenoteNotebookSectionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSiteOnenoteNotebookSectionOperationOptions() CreateSiteOnenoteNotebookSectionOperationOptions {
	return CreateSiteOnenoteNotebookSectionOperationOptions{}
}

func (o CreateSiteOnenoteNotebookSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSiteOnenoteNotebookSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSiteOnenoteNotebookSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSiteOnenoteNotebookSection - Create new navigation property to sections for groups
func (c SiteOnenoteNotebookSectionClient) CreateSiteOnenoteNotebookSection(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookId, input stable.OnenoteSection, options CreateSiteOnenoteNotebookSectionOperationOptions) (result CreateSiteOnenoteNotebookSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sections", id.ID()),
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

	var model stable.OnenoteSection
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
