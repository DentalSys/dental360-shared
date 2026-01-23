package current_user

import (
	"encoding/json"
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

func TestCurrentUser_JSONTags(t *testing.T) {
	t.Run("should return user_id in json", func(t *testing.T) {
		cu := New("user-1", CurrentUserScopeIdentity)

		data, err := json.Marshal(cu)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		if _, ok := result["user_id"]; !ok {
			t.Fatalf("expected key user_id")
		}

		if result["user_id"] != "user-1" {
			t.Fatalf("expected user_id user-1")
		}

		if _, ok := result["UserID"]; ok {
			t.Fatalf("did not expect key UserID")
		}
	})
}
