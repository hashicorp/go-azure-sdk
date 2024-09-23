package grouppolicymigrationreportunsupportedgrouppolicyextension

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

type ListGroupPolicyMigrationReportUnsupportedExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnsupportedGroupPolicyExtension
}

type ListGroupPolicyMigrationReportUnsupportedExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnsupportedGroupPolicyExtension
}

type ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions struct {
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

func DefaultListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions() ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions {
	return ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions{}
}

func (o ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGroupPolicyMigrationReportUnsupportedExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyMigrationReportUnsupportedExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyMigrationReportUnsupportedExtensions - Get unsupportedGroupPolicyExtensions from deviceManagement. A
// list of unsupported group policy extensions inside the Group Policy Object.
func (c GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient) ListGroupPolicyMigrationReportUnsupportedExtensions(ctx context.Context, id beta.DeviceManagementGroupPolicyMigrationReportId, options ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions) (result ListGroupPolicyMigrationReportUnsupportedExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGroupPolicyMigrationReportUnsupportedExtensionsCustomPager{},
		Path:          fmt.Sprintf("%s/unsupportedGroupPolicyExtensions", id.ID()),
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
		Values *[]beta.UnsupportedGroupPolicyExtension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyMigrationReportUnsupportedExtensionsComplete retrieves all the results into a single object
func (c GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient) ListGroupPolicyMigrationReportUnsupportedExtensionsComplete(ctx context.Context, id beta.DeviceManagementGroupPolicyMigrationReportId, options ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions) (ListGroupPolicyMigrationReportUnsupportedExtensionsCompleteResult, error) {
	return c.ListGroupPolicyMigrationReportUnsupportedExtensionsCompleteMatchingPredicate(ctx, id, options, UnsupportedGroupPolicyExtensionOperationPredicate{})
}

// ListGroupPolicyMigrationReportUnsupportedExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient) ListGroupPolicyMigrationReportUnsupportedExtensionsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementGroupPolicyMigrationReportId, options ListGroupPolicyMigrationReportUnsupportedExtensionsOperationOptions, predicate UnsupportedGroupPolicyExtensionOperationPredicate) (result ListGroupPolicyMigrationReportUnsupportedExtensionsCompleteResult, err error) {
	items := make([]beta.UnsupportedGroupPolicyExtension, 0)

	resp, err := c.ListGroupPolicyMigrationReportUnsupportedExtensions(ctx, id, options)
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

	result = ListGroupPolicyMigrationReportUnsupportedExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
