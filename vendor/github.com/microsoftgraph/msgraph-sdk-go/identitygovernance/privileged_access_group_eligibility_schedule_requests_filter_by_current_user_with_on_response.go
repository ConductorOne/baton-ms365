package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse struct {
    PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse
}
// NewPrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse instantiates a new PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse()(*PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse) {
    m := &PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse{
        PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse: *NewPrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreatePrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse(), nil
}
// PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponseable
}
