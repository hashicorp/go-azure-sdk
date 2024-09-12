package todolisttaskattachment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTodoListTaskAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AttachmentBase
}

type ListTodoListTaskAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AttachmentBase
}

type ListTodoListTaskAttachmentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTodoListTaskAttachmentsOperationOptions() ListTodoListTaskAttachmentsOperationOptions {
	return ListTodoListTaskAttachmentsOperationOptions{}
}

func (o ListTodoListTaskAttachmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTodoListTaskAttachmentsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListTodoListTaskAttachmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTodoListTaskAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListTaskAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoListTaskAttachments - List taskFileAttachments. Get a list of the taskFileAttachment objects and their
// properties. The contentBytes property will not be returned in the response. Use the Get attachment API to view the
// contentBytes.
func (c TodoListTaskAttachmentClient) ListTodoListTaskAttachments(ctx context.Context, id stable.MeTodoListIdTaskId, options ListTodoListTaskAttachmentsOperationOptions) (result ListTodoListTaskAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTodoListTaskAttachmentsCustomPager{},
		Path:          fmt.Sprintf("%s/attachments", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.AttachmentBase, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalAttachmentBaseImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.AttachmentBase (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListTodoListTaskAttachmentsComplete retrieves all the results into a single object
func (c TodoListTaskAttachmentClient) ListTodoListTaskAttachmentsComplete(ctx context.Context, id stable.MeTodoListIdTaskId, options ListTodoListTaskAttachmentsOperationOptions) (ListTodoListTaskAttachmentsCompleteResult, error) {
	return c.ListTodoListTaskAttachmentsCompleteMatchingPredicate(ctx, id, options, AttachmentBaseOperationPredicate{})
}

// ListTodoListTaskAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListTaskAttachmentClient) ListTodoListTaskAttachmentsCompleteMatchingPredicate(ctx context.Context, id stable.MeTodoListIdTaskId, options ListTodoListTaskAttachmentsOperationOptions, predicate AttachmentBaseOperationPredicate) (result ListTodoListTaskAttachmentsCompleteResult, err error) {
	items := make([]stable.AttachmentBase, 0)

	resp, err := c.ListTodoListTaskAttachments(ctx, id, options)
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

	result = ListTodoListTaskAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
