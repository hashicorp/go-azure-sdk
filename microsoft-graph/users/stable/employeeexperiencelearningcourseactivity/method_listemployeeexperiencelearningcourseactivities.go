package employeeexperiencelearningcourseactivity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEmployeeExperienceLearningCourseActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.LearningCourseActivity
}

type ListEmployeeExperienceLearningCourseActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.LearningCourseActivity
}

type ListEmployeeExperienceLearningCourseActivitiesOperationOptions struct {
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

// ListEmployeeExperienceLearningCourseActivities - List learningCourseActivities. Get a list of the
// learningCourseActivity objects (assigned or self-initiated) for a user.
func (c EmployeeExperienceLearningCourseActivityClient) ListEmployeeExperienceLearningCourseActivities(ctx context.Context, id stable.UserId, options ListEmployeeExperienceLearningCourseActivitiesOperationOptions) (result ListEmployeeExperienceLearningCourseActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEmployeeExperienceLearningCourseActivitiesCustomPager{},
		Path:          fmt.Sprintf("%s/employeeExperience/learningCourseActivities", id.ID()),
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

	temp := make([]stable.LearningCourseActivity, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalLearningCourseActivityImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.LearningCourseActivity (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListEmployeeExperienceLearningCourseActivitiesComplete retrieves all the results into a single object
func (c EmployeeExperienceLearningCourseActivityClient) ListEmployeeExperienceLearningCourseActivitiesComplete(ctx context.Context, id stable.UserId, options ListEmployeeExperienceLearningCourseActivitiesOperationOptions) (ListEmployeeExperienceLearningCourseActivitiesCompleteResult, error) {
	return c.ListEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate(ctx, id, options, LearningCourseActivityOperationPredicate{})
}

// ListEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EmployeeExperienceLearningCourseActivityClient) ListEmployeeExperienceLearningCourseActivitiesCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListEmployeeExperienceLearningCourseActivitiesOperationOptions, predicate LearningCourseActivityOperationPredicate) (result ListEmployeeExperienceLearningCourseActivitiesCompleteResult, err error) {
	items := make([]stable.LearningCourseActivity, 0)

	resp, err := c.ListEmployeeExperienceLearningCourseActivities(ctx, id, options)
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
