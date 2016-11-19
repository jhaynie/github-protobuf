package github

import (
	"testing"
	"encoding/json"
)

func TestIntegrationInstallationRepositories(t *testing.T) {
	event := `{
	  "action": "removed",
	  "installation": {
	    "id": 2,
	    "account": {
	      "login": "octocat",
	      "id": 1,
	      "avatar_url": "https://github.com/images/error/octocat_happy.gif",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/octocat",
	      "html_url": "https://github.com/octocat",
	      "followers_url": "https://api.github.com/users/octocat/followers",
	      "following_url": "https://api.github.com/users/octocat/following{/other_user}",
	      "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
	      "organizations_url": "https://api.github.com/users/octocat/orgs",
	      "repos_url": "https://api.github.com/users/octocat/repos",
	      "events_url": "https://api.github.com/users/octocat/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/octocat/received_events",
	      "type": "User",
	      "site_admin": false
	    },
	    "access_tokens_url": "https://api.github.com/installations/2/access_tokens",
	    "repositories_url": "https://api.github.com/installation/repositories",
	    "html_url": "https://github.com/settings/installations/2"
	  },
	  "repository_selection": "selected",
	  "repositories_added": [

	  ],
	  "repositories_removed": [
	    {
	      "id": 1296269,
	      "name": "Hello-World",
	      "full_name": "octocat/Hello-World"
	    }
	  ],
	  "sender": {
	    "login": "octocat",
	    "id": 1,
	    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
	    "gravatar_id": "",
	    "url": "https://api.github.com/users/octocat",
	    "html_url": "https://github.com/octocat",
	    "followers_url": "https://api.github.com/users/octocat/followers",
	    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
	    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
	    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
	    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
	    "organizations_url": "https://api.github.com/users/octocat/orgs",
	    "repos_url": "https://api.github.com/users/octocat/repos",
	    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
	    "received_events_url": "https://api.github.com/users/octocat/received_events",
	    "type": "User",
	    "site_admin": false
	  }
	}`
	var integrationEvent IntegrationInstallationRepositories
	err := json.Unmarshal([]byte(event), &integrationEvent)
	if err != nil {
		t.Error(err)
	} else {
		if integrationEvent.Action != "removed" {
			t.Error("integrationEvent.Action != removed")
		}
		if integrationEvent.Sender.Id != 1 {
			t.Error("integrationEvent.Sender.Id != 1")
		}
	}
}
