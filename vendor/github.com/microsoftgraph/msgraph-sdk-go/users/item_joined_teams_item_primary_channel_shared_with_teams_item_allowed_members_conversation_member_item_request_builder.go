package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters
}
// NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal instantiates a new ConversationMemberItemRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) {
    m := &ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/{conversationMember%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder instantiates a new ConversationMemberItemRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a collection of team members who have access to the shared channel.
func (m *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConversationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable), nil
}
// ToGetRequestInformation a collection of team members who have access to the shared channel.
func (m *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) {
    return NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}