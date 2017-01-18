package github

import (
	"testing"
	"encoding/json"
)

func TestUser(t *testing.T) {
	event := `{
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
  }`
	var user User
	err := json.Unmarshal([]byte(event), &user)
	if err != nil {
		t.Error(err)
	} else {
		if user.Login != "octocat" {
			t.Fatalf("expected login to be 'octocat' was '%v'", user.Login)
		}
	}
}

func TestUserWithPermissions(t *testing.T) {
	event := `{
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
    "site_admin": false,
	 "permissions": {
       "pull": true,
       "push": true,
       "admin": false
	  }
  }`
	var user User
	err := json.Unmarshal([]byte(event), &user)
	if err != nil {
		t.Error(err)
	} else {
		if user.Login != "octocat" {
			t.Fatalf("expected login to be 'octocat' was '%v'", user.Login)
		}
		if user.Permissions.Pull != true {
			t.Fatalf("expected permissions.Pull to be 'true' was '%v'", user.Permissions.Pull)
		}
		if user.Permissions.Push != true {
			t.Fatalf("expected permissions.Push to be 'true' was '%v'", user.Permissions.Push)
		}
		if user.Permissions.Admin == true {
			t.Fatalf("expected permissions.Admin to be 'false' was '%v'", user.Permissions.Admin)
		}
	}
}
