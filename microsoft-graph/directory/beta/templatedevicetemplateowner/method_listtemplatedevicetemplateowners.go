package templatedevicetemplateowner

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTemplateDeviceTemplateOwnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListTemplateDeviceTemplateOwnersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListTemplateDeviceTemplateOwnersOperationOptions struct {
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

func DefaultListTemplateDeviceTemplateOwnersOperationOptions() ListTemplateDeviceTemplateOwnersOperationOptions {
	return ListTemplateDeviceTemplateOwnersOperationOptions{}
}

func (o ListTemplateDeviceTemplateOwnersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTemplateDeviceTemplateOwnersOperationOptions) ToOData() *odata.Query {
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

func (o ListTemplateDeviceTemplateOwnersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTemplateDeviceTemplateOwnersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateDeviceTemplateOwnersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateDeviceTemplateOwners - List deviceTemplate owners. Get a list of owners for a deviceTemplate object.
// Owners can be represented as service principals, users, or applications.
func (c TemplateDeviceTemplateOwnerClient) ListTemplateDeviceTemplateOwners(ctx context.Context, id beta.DirectoryTemplateDeviceTemplateId, options ListTemplateDeviceTemplateOwnersOperationOptions) (result ListTemplateDeviceTemplateOwnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTemplateDeviceTemplateOwnersCustomPager{},
		Path:          fmt.Sprintf("%s/owners", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListTemplateDeviceTemplateOwnersComplete retrieves all the results into a single object
func (c TemplateDeviceTemplateOwnerClient) ListTemplateDeviceTemplateOwnersComplete(ctx context.Context, id beta.DirectoryTemplateDeviceTemplateId, options ListTemplateDeviceTemplateOwnersOperationOptions) (ListTemplateDeviceTemplateOwnersCompleteResult, error) {
	return c.ListTemplateDeviceTemplateOwnersCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListTemplateDeviceTemplateOwnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateDeviceTemplateOwnerClient) ListTemplateDeviceTemplateOwnersCompleteMatchingPredicate(ctx context.Context, id beta.DirectoryTemplateDeviceTemplateId, options ListTemplateDeviceTemplateOwnersOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListTemplateDeviceTemplateOwnersCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListTemplateDeviceTemplateOwners(ctx, id, options)
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

	result = ListTemplateDeviceTemplateOwnersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
