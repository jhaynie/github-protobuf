package github

import (
	"testing"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

func TestMessage(t *testing.T) {
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
		buf, err := pingEvent.Marshal()
		if err != nil {
			t.Error(err)
		}
		message := Message{
			Version: 1,
			Timestamp: "2016-11-17T05:21:33Z",
			Id: "123",
			Event: "test",
			Payload: buf,
			Context: "",
		}
		buf, err = proto.Marshal(&message)
		if err != nil {
			t.Error(err)
		}
		var messageT Message
		err = proto.Unmarshal(buf, &messageT)
		if err != nil {
			t.Error(err)
		}
		if messageT.Version != 1 {
			t.Error("messageT.Version != 1")
		}
		var pingEventT Ping
		err = proto.Unmarshal(messageT.Payload, &pingEventT)
		if err != nil {
			t.Error(err)
		}
		if pingEventT.HookId != 10746963 {
			t.Error("pingEventT.HookId != 10746963")
		}
	}
}
