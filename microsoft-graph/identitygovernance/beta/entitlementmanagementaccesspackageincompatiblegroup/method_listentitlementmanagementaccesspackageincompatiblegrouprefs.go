package entitlementmanagementaccesspackageincompatiblegroup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions() ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions {
	return ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageIncompatibleGroupRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageIncompatibleGroupRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageIncompatibleGroupRefs - List incompatibleGroups. Retrieve a list of the group
// objects marked as incompatible on an accessPackage.
func (c EntitlementManagementAccessPackageIncompatibleGroupClient) ListEntitlementManagementAccessPackageIncompatibleGroupRefs(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) (result ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageIncompatibleGroupRefsCustomPager{},
		Path:          fmt.Sprintf("%s/incompatibleGroups/$ref", id.ID()),
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

// ListEntitlementManagementAccessPackageIncompatibleGroupRefsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageIncompatibleGroupClient) ListEntitlementManagementAccessPackageIncompatibleGroupRefsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) (ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageIncompatibleGroupClient) ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListEntitlementManagementAccessPackageIncompatibleGroupRefs(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageIncompatibleGroupRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
