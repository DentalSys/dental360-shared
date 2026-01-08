package currentuser

import (
	"testing"
)

func TestCurrentUser(describe *testing.T) {
	describe.Run("New()", func(test *testing.T) {
		test.Run("should create user with defaults", func(t *testing.T) {
			currentUser := New("user-1", CurrentUserScopeIdentity)

			if currentUser.UserID != "user-1" {
				t.Fatalf("expected userId user-1")
			}
		})
	})
}
