package serviceresource

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceTasksListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProjectTask

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ServiceTasksListOperationResponse, error)
}

type ServiceTasksListCompleteResult struct {
	Items []ProjectTask
}

func (r ServiceTasksListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ServiceTasksListOperationResponse) LoadMore(ctx context.Context) (resp ServiceTasksListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ServiceTasksListOperationOptions struct {
	TaskType *string
}

func DefaultServiceTasksListOperationOptions() ServiceTasksListOperationOptions {
	return ServiceTasksListOperationOptions{}
}

func (o ServiceTasksListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ServiceTasksListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.TaskType != nil {
		out["taskType"] = *o.TaskType
	}

	return out
}

// ServiceTasksList ...
func (c ServiceResourceClient) ServiceTasksList(ctx context.Context, id ServiceId, options ServiceTasksListOperationOptions) (resp ServiceTasksListOperationResponse, err error) {
	req, err := c.preparerForServiceTasksList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serviceresource.ServiceResourceClient", "ServiceTasksList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "serviceresource.ServiceResourceClient", "ServiceTasksList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForServiceTasksList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serviceresource.ServiceResourceClient", "ServiceTasksList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForServiceTasksList prepares the ServiceTasksList request.
func (c ServiceResourceClient) preparerForServiceTasksList(ctx context.Context, id ServiceId, options ServiceTasksListOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/serviceTasks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForServiceTasksListWithNextLink prepares the ServiceTasksList request with the given nextLink token.
func (c ServiceResourceClient) preparerForServiceTasksListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServiceTasksList handles the response to the ServiceTasksList request. The method always
// closes the http.Response Body.
func (c ServiceResourceClient) responderForServiceTasksList(resp *http.Response) (result ServiceTasksListOperationResponse, err error) {
	type page struct {
		Values   []ProjectTask `json:"value"`
		NextLink *string       `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ServiceTasksListOperationResponse, err error) {
			req, err := c.preparerForServiceTasksListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "serviceresource.ServiceResourceClient", "ServiceTasksList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "serviceresource.ServiceResourceClient", "ServiceTasksList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForServiceTasksList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "serviceresource.ServiceResourceClient", "ServiceTasksList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ServiceTasksListComplete retrieves all of the results into a single object
func (c ServiceResourceClient) ServiceTasksListComplete(ctx context.Context, id ServiceId, options ServiceTasksListOperationOptions) (ServiceTasksListCompleteResult, error) {
	return c.ServiceTasksListCompleteMatchingPredicate(ctx, id, options, ProjectTaskOperationPredicate{})
}

// ServiceTasksListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ServiceResourceClient) ServiceTasksListCompleteMatchingPredicate(ctx context.Context, id ServiceId, options ServiceTasksListOperationOptions, predicate ProjectTaskOperationPredicate) (resp ServiceTasksListCompleteResult, err error) {
	items := make([]ProjectTask, 0)

	page, err := c.ServiceTasksList(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ServiceTasksListCompleteResult{
		Items: items,
	}
	return out, nil
}
