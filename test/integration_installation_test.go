package github

import (
	"testing"
	"encoding/json"
)

func TestIntegrationInstallation(t *testing.T) {
	event := `{
	  "action": "created",
	  "installation": {
	    "id": 3171,
	    "account": {
	      "login": "EmpathyAI",
	      "id": 6749695,
	      "avatar_url": "https://avatars.githubusercontent.com/u/6749695?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/EmpathyAI",
	      "html_url": "https://github.com/EmpathyAI",
	      "followers_url": "https://api.github.com/users/EmpathyAI/followers",
	      "following_url": "https://api.github.com/users/EmpathyAI/following{/other_user}",
	      "gists_url": "https://api.github.com/users/EmpathyAI/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/EmpathyAI/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/EmpathyAI/subscriptions",
	      "organizations_url": "https://api.github.com/users/EmpathyAI/orgs",
	      "repos_url": "https://api.github.com/users/EmpathyAI/repos",
	      "events_url": "https://api.github.com/users/EmpathyAI/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/EmpathyAI/received_events",
	      "type": "Organization",
	      "site_admin": false
	    },
	    "access_tokens_url": "https://api.github.com/installations/3171/access_tokens",
	    "repositories_url": "https://api.github.com/installation/repositories",
	    "html_url": "https://github.com/organizations/EmpathyAI/settings/installations/3171"
	  },
	  "sender": {
	    "login": "jhaynie",
	    "id": 6027,
	    "avatar_url": "https://avatars.githubusercontent.com/u/6027?v=3",
	    "gravatar_id": "",
	    "url": "https://api.github.com/users/jhaynie",
	    "html_url": "https://github.com/jhaynie",
	    "followers_url": "https://api.github.com/users/jhaynie/followers",
	    "following_url": "https://api.github.com/users/jhaynie/following{/other_user}",
	    "gists_url": "https://api.github.com/users/jhaynie/gists{/gist_id}",
	    "starred_url": "https://api.github.com/users/jhaynie/starred{/owner}{/repo}",
	    "subscriptions_url": "https://api.github.com/users/jhaynie/subscriptions",
	    "organizations_url": "https://api.github.com/users/jhaynie/orgs",
	    "repos_url": "https://api.github.com/users/jhaynie/repos",
	    "events_url": "https://api.github.com/users/jhaynie/events{/privacy}",
	    "received_events_url": "https://api.github.com/users/jhaynie/received_events",
	    "type": "User",
	    "site_admin": false
	  }
	}`
	var integrationEvent IntegrationInstallation
	err := json.Unmarshal([]byte(event), &integrationEvent)
	if err != nil {
		t.Error(err)
	} else {
		if integrationEvent.Action != "created" {
			t.Error("integrationEvent.Action != created")
		}
		if integrationEvent.Sender.Id != 6027 {
			t.Error("integrationEvent.Sender.Id != 6027")
		}
	}
}
