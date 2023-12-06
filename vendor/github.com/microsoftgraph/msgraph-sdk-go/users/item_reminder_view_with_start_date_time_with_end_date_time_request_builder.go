package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the reminderView method.
type ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters invoke function reminderView
type ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *string, startDateTime *string)(*ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/reminderView(StartDateTime='{StartDateTime}',EndDateTime='{EndDateTime}'){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["EndDateTime"] = *endDateTime
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["StartDateTime"] = *startDateTime
    }
    return m
}
// NewItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function reminderView
// Deprecated: This method is obsolete. Use GetAsReminderViewWithStartDateTimeWithEndDateTimeGetResponse instead.
func (m *ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ItemReminderViewWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemReminderViewWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemReminderViewWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsReminderViewWithStartDateTimeWithEndDateTimeGetResponse invoke function reminderView
func (m *ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsReminderViewWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ItemReminderViewWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemReminderViewWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemReminderViewWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation invoke function reminderView
func (m *ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*ItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
