package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse struct {
    EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponse
}
// NewEntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse instantiates a new EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse and sets the default values.
func NewEntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse()(*EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse) {
    m := &EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse{
        EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponse: *NewEntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse(), nil
}
// EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponseable interface {
    EntitlementManagementRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
