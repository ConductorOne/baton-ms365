package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponse struct {
    ItemJoinedTeamsItemPrimaryChannelMessagesDeltaGetResponse
}
// NewItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponse instantiates a new ItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponse and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponse()(*ItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponse) {
    m := &ItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponse{
        ItemJoinedTeamsItemPrimaryChannelMessagesDeltaGetResponse: *NewItemJoinedTeamsItemPrimaryChannelMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponse(), nil
}
// ItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemJoinedTeamsItemPrimaryChannelMessagesDeltaResponseable interface {
    ItemJoinedTeamsItemPrimaryChannelMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
