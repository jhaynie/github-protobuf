package github

import (
	"testing"
	"encoding/json"
)

func TestPingEvent(t *testing.T) {
	event := `{
	"zen": "Non-blocking is better than blocking.",
	"hook_id": 10746963,
	"hook": {
	  "type": "Integration",
	  "id": 10746963,
	  "name": "web",
	  "active": true,
	  "events": [
		 "commit_comment",
		 "create",
		 "delete",
		 "deployment",
		 "deployment_status",
		 "fork",
		 "issues",
		 "issue_comment",
		 "membership",
		 "public",
		 "pull_request",
		 "pull_request_review",
		 "pull_request_review_comment",
		 "push",
		 "release",
		 "repository",
		 "status",
		 "team_add",
		 "watch"
	  ],
	  "config": {
		 "content_type": "json",
		 "insecure_ssl": "0",
		 "url": "http://61464913.ngrok.io/"
	  },
	  "updated_at": "2016-11-17T05:21:33Z",
	  "created_at": "2016-11-17T05:21:33Z",
	  "integration_id": 661
	}
	}`
	var pingEvent Ping
	err := json.Unmarshal([]byte(event), &pingEvent)
	if err != nil {
		t.Error(err)
	} else {
		if pingEvent.HookId != 10746963 {
			t.Error("pingEvent.HookId != 10746963")
		}
		if pingEvent.Hook.IntegrationId != 661 {
			t.Error("pingEvent.Hook.IntegrationId != 661")
		}
	}
}
