package rolemanagementalertincident

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

type ListRoleManagementAlertIncidentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleManagementAlertIncident
}

type ListRoleManagementAlertIncidentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleManagementAlertIncident
}

type ListRoleManagementAlertIncidentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListRoleManagementAlertIncidentsOperationOptions() ListRoleManagementAlertIncidentsOperationOptions {
	return ListRoleManagementAlertIncidentsOperationOptions{}
}

func (o ListRoleManagementAlertIncidentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleManagementAlertIncidentsOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleManagementAlertIncidentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleManagementAlertIncidentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementAlertIncidentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementAlertIncidents - Get alertIncidents from identityGovernance. Represents the incidents of this type
// of alert that have been triggered in Privileged Identity Management (PIM) for Microsoft Entra roles in the tenant.
// Supports $expand.
func (c RoleManagementAlertIncidentClient) ListRoleManagementAlertIncidents(ctx context.Context, id beta.IdentityGovernanceRoleManagementAlertAlertId, options ListRoleManagementAlertIncidentsOperationOptions) (result ListRoleManagementAlertIncidentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleManagementAlertIncidentsCustomPager{},
		Path:          fmt.Sprintf("%s/alertIncidents", id.ID()),
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

	temp := make([]beta.UnifiedRoleManagementAlertIncident, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalUnifiedRoleManagementAlertIncidentImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.UnifiedRoleManagementAlertIncident (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = temp

	return
}

// ListRoleManagementAlertIncidentsComplete retrieves all the results into a single object
func (c RoleManagementAlertIncidentClient) ListRoleManagementAlertIncidentsComplete(ctx context.Context, id beta.IdentityGovernanceRoleManagementAlertAlertId, options ListRoleManagementAlertIncidentsOperationOptions) (ListRoleManagementAlertIncidentsCompleteResult, error) {
	return c.ListRoleManagementAlertIncidentsCompleteMatchingPredicate(ctx, id, options, UnifiedRoleManagementAlertIncidentOperationPredicate{})
}

// ListRoleManagementAlertIncidentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementAlertIncidentClient) ListRoleManagementAlertIncidentsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceRoleManagementAlertAlertId, options ListRoleManagementAlertIncidentsOperationOptions, predicate UnifiedRoleManagementAlertIncidentOperationPredicate) (result ListRoleManagementAlertIncidentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleManagementAlertIncident, 0)

	resp, err := c.ListRoleManagementAlertIncidents(ctx, id, options)
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

	result = ListRoleManagementAlertIncidentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
