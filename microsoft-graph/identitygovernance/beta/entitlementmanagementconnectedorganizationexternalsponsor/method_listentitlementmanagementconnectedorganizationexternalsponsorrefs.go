package entitlementmanagementconnectedorganizationexternalsponsor

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

type ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions struct {
	Count   *bool
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions() ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions {
	return ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions{}
}

func (o ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementConnectedOrganizationExternalSponsorRefs - List externalSponsors. Retrieve a list of a
// connectedOrganization's external sponsors. The external sponsors are a set of users who can approve requests on
// behalf of other users from that connected organization.
func (c EntitlementManagementConnectedOrganizationExternalSponsorClient) ListEntitlementManagementConnectedOrganizationExternalSponsorRefs(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementConnectedOrganizationId, options ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) (result ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCustomPager{},
		Path:          fmt.Sprintf("%s/externalSponsors/$ref", id.ID()),
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

// ListEntitlementManagementConnectedOrganizationExternalSponsorRefsComplete retrieves all the results into a single object
func (c EntitlementManagementConnectedOrganizationExternalSponsorClient) ListEntitlementManagementConnectedOrganizationExternalSponsorRefsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementConnectedOrganizationId, options ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) (ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCompleteResult, error) {
	return c.ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementConnectedOrganizationExternalSponsorClient) ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementConnectedOrganizationId, options ListEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListEntitlementManagementConnectedOrganizationExternalSponsorRefs(ctx, id, options)
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

	result = ListEntitlementManagementConnectedOrganizationExternalSponsorRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
