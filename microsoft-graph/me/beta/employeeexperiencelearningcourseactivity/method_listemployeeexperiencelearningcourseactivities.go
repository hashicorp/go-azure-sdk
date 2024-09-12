package employeeexperiencelearningcourseactivity

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

type ListEmployeeExperienceLearningCourseActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.LearningCourseActivity
}

type ListEmployeeExperienceLearningCourseActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.LearningCourseActivity
}

type ListEmployeeExperienceLearningCourseActivitiesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEmployeeExperienceLearningCourseActivitiesOperationOptions() ListEmployeeExperienceLearningCourseActivitiesOperationOptions {
	return ListEmployeeExperienceLearningCourseActivitiesOperationOptions{}
}

func (o ListEmployeeExperienceLearningCourseActivitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEmployeeExperienceLearningCourseActivitiesOperationOptions) ToOData() *odata.Query {
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

func (o ListEmployeeExperienceLearningCourseActivitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEmployeeExperienceLearningCourseActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEmployeeExperienceLearningCourseActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEmployeeExperienceLearningCourseActivities - Get learningCourseActivities from me
func (c EmployeeExperienceLearningCourseActivityClient) ListEmployeeExperienceLearningCourseActivities(ctx context.Context, options ListEmployeeExperienceLearningCourseActivitiesOperationOptions) (result ListEmployeeExperienceLearningCourseActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEmployeeExperienceLearningCourseActivitiesCustomPager{},
		Path:          "/me/employeeExperience/learningCourseActivities",
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

	temp := make([]beta.LearningCourseActivity, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalLearningCourseActivityImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.LearningCourseActivity (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListEmployeeExperienceLearningCourseActivitiesComplete retrieves all the results into a single object
func (c EmployeeExperienceLearningCourseActivityClient) ListEmployeeExperienceLearningCourseActivitiesComplete(ctx context.Context, options ListEmployeeExperienceLearningCourseActivitiesOperationOptions) (ListEmployeeExperienceLearningCourseActivitiesCompleteResult, error) {
	return c.ListEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate(ctx, options, LearningCourseActivityOperationPredicate{})
}

// ListEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EmployeeExperienceLearningCourseActivityClient) ListEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate(ctx context.Context, options ListEmployeeExperienceLearningCourseActivitiesOperationOptions, predicate LearningCourseActivityOperationPredicate) (result ListEmployeeExperienceLearningCourseActivitiesCompleteResult, err error) {
	items := make([]beta.LearningCourseActivity, 0)

	resp, err := c.ListEmployeeExperienceLearningCourseActivities(ctx, options)
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

	result = ListEmployeeExperienceLearningCourseActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
