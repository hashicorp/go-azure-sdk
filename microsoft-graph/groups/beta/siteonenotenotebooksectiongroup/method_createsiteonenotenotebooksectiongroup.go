package siteonenotenotebooksectiongroup

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

type CreateSiteOnenoteNotebookSectionGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SectionGroup
}

type CreateSiteOnenoteNotebookSectionGroupOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateSiteOnenoteNotebookSectionGroupOperationOptions() CreateSiteOnenoteNotebookSectionGroupOperationOptions {
	return CreateSiteOnenoteNotebookSectionGroupOperationOptions{}
}

func (o CreateSiteOnenoteNotebookSectionGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSiteOnenoteNotebookSectionGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSiteOnenoteNotebookSectionGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSiteOnenoteNotebookSectionGroup - Create new navigation property to sectionGroups for groups
func (c SiteOnenoteNotebookSectionGroupClient) CreateSiteOnenoteNotebookSectionGroup(ctx context.Context, id beta.GroupIdSiteIdOnenoteNotebookId, input beta.SectionGroup, options CreateSiteOnenoteNotebookSectionGroupOperationOptions) (result CreateSiteOnenoteNotebookSectionGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sectionGroups", id.ID()),
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

	var model beta.SectionGroup
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
