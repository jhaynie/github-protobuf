package github

import (
	"testing"
	"encoding/json"
)

func TestReaction(t *testing.T) {
	event := `{
    "id": 1,
    "user": {
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
    "content": "heart",
    "created_at": "2016-05-20T20:09:31Z"
  }`
  var reaction Reaction
  err := json.Unmarshal([]byte(event), &reaction)
  if err != nil {
	  t.Error(err)
  } else {
	  if reaction.Id != 1 {
		  t.Fatalf("expected Id to be '1', was '%d'", reaction.Id)
	  }
	  if reaction.User.Id != 1 {
		  t.Fatalf("expected User.Id to be '1', was '%d'", reaction.User.Id)
	  }
	  if reaction.Content != "heart" {
		  t.Fatalf("expected Content to be 'heart', was '%s'", reaction.Content)
	  }
	  if reaction.CreatedAt != "2016-05-20T20:09:31Z" {
		  t.Fatalf("expected CreatedAt to be '2016-05-20T20:09:31Z', was '%s'", reaction.CreatedAt)
	  }
  }
}
