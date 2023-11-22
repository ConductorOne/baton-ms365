package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamPrimaryChannelMembersAddResponse 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemTeamPrimaryChannelMembersAddResponse struct {
    ItemTeamPrimaryChannelMembersAddPostResponse
}
// NewItemTeamPrimaryChannelMembersAddResponse instantiates a new ItemTeamPrimaryChannelMembersAddResponse and sets the default values.
func NewItemTeamPrimaryChannelMembersAddResponse()(*ItemTeamPrimaryChannelMembersAddResponse) {
    m := &ItemTeamPrimaryChannelMembersAddResponse{
        ItemTeamPrimaryChannelMembersAddPostResponse: *NewItemTeamPrimaryChannelMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamPrimaryChannelMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamPrimaryChannelMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimaryChannelMembersAddResponse(), nil
}
// ItemTeamPrimaryChannelMembersAddResponseable 
// Deprecated: This class is obsolete. Use addPostResponse instead.
type ItemTeamPrimaryChannelMembersAddResponseable interface {
    ItemTeamPrimaryChannelMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
