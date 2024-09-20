package appleuserinitiatedenrollmentprofileassignment

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

type ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppleEnrollmentProfileAssignment
}

type ListAppleUserInitiatedEnrollmentProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppleEnrollmentProfileAssignment
}

type ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions() ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions {
	return ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions{}
}

func (o ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppleUserInitiatedEnrollmentProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppleUserInitiatedEnrollmentProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppleUserInitiatedEnrollmentProfileAssignments - Get assignments from deviceManagement. The list of assignments
// for this profile.
func (c AppleUserInitiatedEnrollmentProfileAssignmentClient) ListAppleUserInitiatedEnrollmentProfileAssignments(ctx context.Context, id beta.DeviceManagementAppleUserInitiatedEnrollmentProfileId, options ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions) (result ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppleUserInitiatedEnrollmentProfileAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.AppleEnrollmentProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppleUserInitiatedEnrollmentProfileAssignmentsComplete retrieves all the results into a single object
func (c AppleUserInitiatedEnrollmentProfileAssignmentClient) ListAppleUserInitiatedEnrollmentProfileAssignmentsComplete(ctx context.Context, id beta.DeviceManagementAppleUserInitiatedEnrollmentProfileId, options ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions) (ListAppleUserInitiatedEnrollmentProfileAssignmentsCompleteResult, error) {
	return c.ListAppleUserInitiatedEnrollmentProfileAssignmentsCompleteMatchingPredicate(ctx, id, options, AppleEnrollmentProfileAssignmentOperationPredicate{})
}

// ListAppleUserInitiatedEnrollmentProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppleUserInitiatedEnrollmentProfileAssignmentClient) ListAppleUserInitiatedEnrollmentProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementAppleUserInitiatedEnrollmentProfileId, options ListAppleUserInitiatedEnrollmentProfileAssignmentsOperationOptions, predicate AppleEnrollmentProfileAssignmentOperationPredicate) (result ListAppleUserInitiatedEnrollmentProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.AppleEnrollmentProfileAssignment, 0)

	resp, err := c.ListAppleUserInitiatedEnrollmentProfileAssignments(ctx, id, options)
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

	result = ListAppleUserInitiatedEnrollmentProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
