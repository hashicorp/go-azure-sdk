package diagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsDiagnosticsListRevisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Revision
}

type ContainerAppsDiagnosticsListRevisionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Revision
}

type ContainerAppsDiagnosticsListRevisionsOperationOptions struct {
	Filter *string
}

func DefaultContainerAppsDiagnosticsListRevisionsOperationOptions() ContainerAppsDiagnosticsListRevisionsOperationOptions {
	return ContainerAppsDiagnosticsListRevisionsOperationOptions{}
}

func (o ContainerAppsDiagnosticsListRevisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ContainerAppsDiagnosticsListRevisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ContainerAppsDiagnosticsListRevisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ContainerAppsDiagnosticsListRevisions ...
func (c DiagnosticsClient) ContainerAppsDiagnosticsListRevisions(ctx context.Context, id ContainerAppId, options ContainerAppsDiagnosticsListRevisionsOperationOptions) (result ContainerAppsDiagnosticsListRevisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/detectorProperties/revisionsApi/revisions", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]Revision `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ContainerAppsDiagnosticsListRevisionsComplete retrieves all the results into a single object
func (c DiagnosticsClient) ContainerAppsDiagnosticsListRevisionsComplete(ctx context.Context, id ContainerAppId, options ContainerAppsDiagnosticsListRevisionsOperationOptions) (ContainerAppsDiagnosticsListRevisionsCompleteResult, error) {
	return c.ContainerAppsDiagnosticsListRevisionsCompleteMatchingPredicate(ctx, id, options, RevisionOperationPredicate{})
}

// ContainerAppsDiagnosticsListRevisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ContainerAppsDiagnosticsListRevisionsCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, options ContainerAppsDiagnosticsListRevisionsOperationOptions, predicate RevisionOperationPredicate) (result ContainerAppsDiagnosticsListRevisionsCompleteResult, err error) {
	items := make([]Revision, 0)

	resp, err := c.ContainerAppsDiagnosticsListRevisions(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ContainerAppsDiagnosticsListRevisionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
