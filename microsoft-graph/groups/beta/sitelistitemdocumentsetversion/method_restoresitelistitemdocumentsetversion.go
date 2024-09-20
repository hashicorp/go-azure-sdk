package sitelistitemdocumentsetversion

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

type RestoreSiteListItemDocumentSetVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RestoreSiteListItemDocumentSetVersionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultRestoreSiteListItemDocumentSetVersionOperationOptions() RestoreSiteListItemDocumentSetVersionOperationOptions {
	return RestoreSiteListItemDocumentSetVersionOperationOptions{}
}

func (o RestoreSiteListItemDocumentSetVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestoreSiteListItemDocumentSetVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RestoreSiteListItemDocumentSetVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RestoreSiteListItemDocumentSetVersion - Invoke action restore. Restore a document set version.
func (c SiteListItemDocumentSetVersionClient) RestoreSiteListItemDocumentSetVersion(ctx context.Context, id beta.GroupIdSiteIdListIdItemIdDocumentSetVersionId, options RestoreSiteListItemDocumentSetVersionOperationOptions) (result RestoreSiteListItemDocumentSetVersionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/restore", id.ID()),
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
