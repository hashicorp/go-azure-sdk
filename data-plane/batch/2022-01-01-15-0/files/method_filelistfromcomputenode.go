package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileListFromComputeNodeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NodeFile
}

type FileListFromComputeNodeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NodeFile
}

type FileListFromComputeNodeOperationOptions struct {
	ClientRequestId       *string
	Filter                *string
	Maxresults            *int64
	OcpDate               *string
	Recursive             *bool
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultFileListFromComputeNodeOperationOptions() FileListFromComputeNodeOperationOptions {
	return FileListFromComputeNodeOperationOptions{}
}

func (o FileListFromComputeNodeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.ClientRequestId != nil {
		out.Append("client-request-id", fmt.Sprintf("%v", *o.ClientRequestId))
	}
	if o.OcpDate != nil {
		out.Append("ocp-date", fmt.Sprintf("%v", *o.OcpDate))
	}
	if o.ReturnClientRequestId != nil {
		out.Append("return-client-request-id", fmt.Sprintf("%v", *o.ReturnClientRequestId))
	}
	return &out
}

func (o FileListFromComputeNodeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o FileListFromComputeNodeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	if o.Recursive != nil {
		out.Append("recursive", fmt.Sprintf("%v", *o.Recursive))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

type FileListFromComputeNodeCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *FileListFromComputeNodeCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// FileListFromComputeNode ...
func (c FilesClient) FileListFromComputeNode(ctx context.Context, id NodeId, options FileListFromComputeNodeOperationOptions) (result FileListFromComputeNodeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &FileListFromComputeNodeCustomPager{},
		Path:          fmt.Sprintf("%s/files", id.ID()),
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
		Values *[]NodeFile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// FileListFromComputeNodeComplete retrieves all the results into a single object
func (c FilesClient) FileListFromComputeNodeComplete(ctx context.Context, id NodeId, options FileListFromComputeNodeOperationOptions) (FileListFromComputeNodeCompleteResult, error) {
	return c.FileListFromComputeNodeCompleteMatchingPredicate(ctx, id, options, NodeFileOperationPredicate{})
}

// FileListFromComputeNodeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FilesClient) FileListFromComputeNodeCompleteMatchingPredicate(ctx context.Context, id NodeId, options FileListFromComputeNodeOperationOptions, predicate NodeFileOperationPredicate) (result FileListFromComputeNodeCompleteResult, err error) {
	items := make([]NodeFile, 0)

	resp, err := c.FileListFromComputeNode(ctx, id, options)
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

	result = FileListFromComputeNodeCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
