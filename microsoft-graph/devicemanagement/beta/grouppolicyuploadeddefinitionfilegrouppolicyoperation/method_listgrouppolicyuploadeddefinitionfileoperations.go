package grouppolicyuploadeddefinitionfilegrouppolicyoperation

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

type ListGroupPolicyUploadedDefinitionFileOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyOperation
}

type ListGroupPolicyUploadedDefinitionFileOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyOperation
}

type ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListGroupPolicyUploadedDefinitionFileOperationsOperationOptions() ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions {
	return ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions{}
}

func (o ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGroupPolicyUploadedDefinitionFileOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyUploadedDefinitionFileOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyUploadedDefinitionFileOperations - Get groupPolicyOperations from deviceManagement. The list of
// operations on the uploaded ADMX file.
func (c GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient) ListGroupPolicyUploadedDefinitionFileOperations(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, options ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions) (result ListGroupPolicyUploadedDefinitionFileOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGroupPolicyUploadedDefinitionFileOperationsCustomPager{},
		Path:          fmt.Sprintf("%s/groupPolicyOperations", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.GroupPolicyOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyUploadedDefinitionFileOperationsComplete retrieves all the results into a single object
func (c GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient) ListGroupPolicyUploadedDefinitionFileOperationsComplete(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, options ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions) (ListGroupPolicyUploadedDefinitionFileOperationsCompleteResult, error) {
	return c.ListGroupPolicyUploadedDefinitionFileOperationsCompleteMatchingPredicate(ctx, id, options, GroupPolicyOperationOperationPredicate{})
}

// ListGroupPolicyUploadedDefinitionFileOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient) ListGroupPolicyUploadedDefinitionFileOperationsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementGroupPolicyUploadedDefinitionFileId, options ListGroupPolicyUploadedDefinitionFileOperationsOperationOptions, predicate GroupPolicyOperationOperationPredicate) (result ListGroupPolicyUploadedDefinitionFileOperationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyOperation, 0)

	resp, err := c.ListGroupPolicyUploadedDefinitionFileOperations(ctx, id, options)
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

	result = ListGroupPolicyUploadedDefinitionFileOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
