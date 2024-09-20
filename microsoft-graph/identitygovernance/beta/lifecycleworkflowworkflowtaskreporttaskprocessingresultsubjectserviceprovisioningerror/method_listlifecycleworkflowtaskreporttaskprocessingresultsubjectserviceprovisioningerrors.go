package lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror

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

type ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServiceProvisioningError
}

type ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServiceProvisioningError
}

type ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions struct {
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

func DefaultListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions() ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions {
	return ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions{}
}

func (o ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrors - Get serviceProvisioningErrors
// property value. Errors published by a federated service describing a nontransient, service-specific error regarding
// the properties or link from a user object.
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient) ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrors(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId, options ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions) (result ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCustomPager{},
		Path:          fmt.Sprintf("%s/subject/serviceProvisioningErrors", id.ID()),
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

	temp := make([]beta.ServiceProvisioningError, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalServiceProvisioningErrorImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ServiceProvisioningError (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient) ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsComplete(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId, options ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions) (ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteResult, error) {
	return c.ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, options, ServiceProvisioningErrorOperationPredicate{})
}

// ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient) ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId, options ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsOperationOptions, predicate ServiceProvisioningErrorOperationPredicate) (result ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]beta.ServiceProvisioningError, 0)

	resp, err := c.ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrors(ctx, id, options)
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

	result = ListLifecycleWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
