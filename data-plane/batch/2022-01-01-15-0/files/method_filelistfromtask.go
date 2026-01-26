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

type FileListFromTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NodeFile
}

type FileListFromTaskCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NodeFile
}

type FileListFromTaskOperationOptions struct {
	ClientRequestId       *string
	Filter                *string
	Maxresults            *int64
	OcpDate               *string
	Recursive             *bool
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultFileListFromTaskOperationOptions() FileListFromTaskOperationOptions {
	return FileListFromTaskOperationOptions{}
}

func (o FileListFromTaskOperationOptions) ToHeaders() *client.Headers {
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

func (o FileListFromTaskOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o FileListFromTaskOperationOptions) ToQuery() *client.QueryParams {
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

type FileListFromTaskCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *FileListFromTaskCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// FileListFromTask ...
func (c FilesClient) FileListFromTask(ctx context.Context, id TaskId, options FileListFromTaskOperationOptions) (result FileListFromTaskOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &FileListFromTaskCustomPager{},
		Path:          fmt.Sprintf("%s/files", id.Path()),
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

// FileListFromTaskComplete retrieves all the results into a single object
func (c FilesClient) FileListFromTaskComplete(ctx context.Context, id TaskId, options FileListFromTaskOperationOptions) (FileListFromTaskCompleteResult, error) {
	return c.FileListFromTaskCompleteMatchingPredicate(ctx, id, options, NodeFileOperationPredicate{})
}

// FileListFromTaskCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FilesClient) FileListFromTaskCompleteMatchingPredicate(ctx context.Context, id TaskId, options FileListFromTaskOperationOptions, predicate NodeFileOperationPredicate) (result FileListFromTaskCompleteResult, err error) {
	items := make([]NodeFile, 0)

	resp, err := c.FileListFromTask(ctx, id, options)
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

	result = FileListFromTaskCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
