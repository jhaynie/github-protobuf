package github

import (
	"testing"
	"encoding/json"
)

func TestLabelEvent(t *testing.T) {
	event := `{
	  "action": "edited",
	  "label": {
	    "id": 574435088,
	    "url": "https://api.github.com/repos/pinpt/test_nodejs/labels/foobar2",
	    "name": "foobar2",
	    "color": "5319e7",
	    "default": false
	  },
	  "changes": {
	    "name": {
	      "from": "foobar"
	    }
	  },
	  "repository": {
	    "id": 85026034,
	    "name": "test_nodejs",
	    "full_name": "pinpt/test_nodejs",
	    "owner": {
	      "login": "pinpt",
	      "id": 24400526,
	      "avatar_url": "https://avatars0.githubusercontent.com/u/24400526?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/pinpt",
	      "html_url": "https://github.com/pinpt",
	      "followers_url": "https://api.github.com/users/pinpt/followers",
	      "following_url": "https://api.github.com/users/pinpt/following{/other_user}",
	      "gists_url": "https://api.github.com/users/pinpt/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/pinpt/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/pinpt/subscriptions",
	      "organizations_url": "https://api.github.com/users/pinpt/orgs",
	      "repos_url": "https://api.github.com/users/pinpt/repos",
	      "events_url": "https://api.github.com/users/pinpt/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/pinpt/received_events",
	      "type": "Organization",
	      "site_admin": false
	    },
	    "private": false,
	    "html_url": "https://github.com/pinpt/test_nodejs",
	    "description": "Test NodeJS Project",
	    "fork": false,
	    "url": "https://api.github.com/repos/pinpt/test_nodejs",
	    "forks_url": "https://api.github.com/repos/pinpt/test_nodejs/forks",
	    "keys_url": "https://api.github.com/repos/pinpt/test_nodejs/keys{/key_id}",
	    "collaborators_url": "https://api.github.com/repos/pinpt/test_nodejs/collaborators{/collaborator}",
	    "teams_url": "https://api.github.com/repos/pinpt/test_nodejs/teams",
	    "hooks_url": "https://api.github.com/repos/pinpt/test_nodejs/hooks",
	    "issue_events_url": "https://api.github.com/repos/pinpt/test_nodejs/issues/events{/number}",
	    "events_url": "https://api.github.com/repos/pinpt/test_nodejs/events",
	    "assignees_url": "https://api.github.com/repos/pinpt/test_nodejs/assignees{/user}",
	    "branches_url": "https://api.github.com/repos/pinpt/test_nodejs/branches{/branch}",
	    "tags_url": "https://api.github.com/repos/pinpt/test_nodejs/tags",
	    "blobs_url": "https://api.github.com/repos/pinpt/test_nodejs/git/blobs{/sha}",
	    "git_tags_url": "https://api.github.com/repos/pinpt/test_nodejs/git/tags{/sha}",
	    "git_refs_url": "https://api.github.com/repos/pinpt/test_nodejs/git/refs{/sha}",
	    "trees_url": "https://api.github.com/repos/pinpt/test_nodejs/git/trees{/sha}",
	    "statuses_url": "https://api.github.com/repos/pinpt/test_nodejs/statuses/{sha}",
	    "languages_url": "https://api.github.com/repos/pinpt/test_nodejs/languages",
	    "stargazers_url": "https://api.github.com/repos/pinpt/test_nodejs/stargazers",
	    "contributors_url": "https://api.github.com/repos/pinpt/test_nodejs/contributors",
	    "subscribers_url": "https://api.github.com/repos/pinpt/test_nodejs/subscribers",
	    "subscription_url": "https://api.github.com/repos/pinpt/test_nodejs/subscription",
	    "commits_url": "https://api.github.com/repos/pinpt/test_nodejs/commits{/sha}",
	    "git_commits_url": "https://api.github.com/repos/pinpt/test_nodejs/git/commits{/sha}",
	    "comments_url": "https://api.github.com/repos/pinpt/test_nodejs/comments{/number}",
	    "issue_comment_url": "https://api.github.com/repos/pinpt/test_nodejs/issues/comments{/number}",
	    "contents_url": "https://api.github.com/repos/pinpt/test_nodejs/contents/{+path}",
	    "compare_url": "https://api.github.com/repos/pinpt/test_nodejs/compare/{base}...{head}",
	    "merges_url": "https://api.github.com/repos/pinpt/test_nodejs/merges",
	    "archive_url": "https://api.github.com/repos/pinpt/test_nodejs/{archive_format}{/ref}",
	    "downloads_url": "https://api.github.com/repos/pinpt/test_nodejs/downloads",
	    "issues_url": "https://api.github.com/repos/pinpt/test_nodejs/issues{/number}",
	    "pulls_url": "https://api.github.com/repos/pinpt/test_nodejs/pulls{/number}",
	    "milestones_url": "https://api.github.com/repos/pinpt/test_nodejs/milestones{/number}",
	    "notifications_url": "https://api.github.com/repos/pinpt/test_nodejs/notifications{?since,all,participating}",
	    "labels_url": "https://api.github.com/repos/pinpt/test_nodejs/labels{/name}",
	    "releases_url": "https://api.github.com/repos/pinpt/test_nodejs/releases{/id}",
	    "deployments_url": "https://api.github.com/repos/pinpt/test_nodejs/deployments",
	    "created_at": "2017-03-15T03:41:38Z",
	    "updated_at": "2017-03-15T03:47:43Z",
	    "pushed_at": "2017-03-16T23:32:55Z",
	    "git_url": "git://github.com/pinpt/test_nodejs.git",
	    "ssh_url": "git@github.com:pinpt/test_nodejs.git",
	    "clone_url": "https://github.com/pinpt/test_nodejs.git",
	    "svn_url": "https://github.com/pinpt/test_nodejs",
	    "homepage": null,
	    "size": 34,
	    "stargazers_count": 0,
	    "watchers_count": 0,
	    "language": "JavaScript",
	    "has_issues": true,
	    "has_projects": true,
	    "has_downloads": true,
	    "has_wiki": true,
	    "has_pages": false,
	    "forks_count": 0,
	    "mirror_url": null,
	    "open_issues_count": 0,
	    "forks": 0,
	    "open_issues": 0,
	    "watchers": 0,
	    "default_branch": "master"
	  },
	  "organization": {
	    "login": "pinpt",
	    "id": 24400526,
	    "url": "https://api.github.com/orgs/pinpt",
	    "repos_url": "https://api.github.com/orgs/pinpt/repos",
	    "events_url": "https://api.github.com/orgs/pinpt/events",
	    "hooks_url": "https://api.github.com/orgs/pinpt/hooks",
	    "issues_url": "https://api.github.com/orgs/pinpt/issues",
	    "members_url": "https://api.github.com/orgs/pinpt/members{/member}",
	    "public_members_url": "https://api.github.com/orgs/pinpt/public_members{/member}",
	    "avatar_url": "https://avatars0.githubusercontent.com/u/24400526?v=3",
	    "description": ""
	  },
	  "sender": {
	    "login": "jhaynie",
	    "id": 6027,
	    "avatar_url": "https://avatars2.githubusercontent.com/u/6027?v=3",
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
	  },
	  "installation": {
	    "id": 15899
	  }
	}`
	var labelEvent LabelEvent
	err := json.Unmarshal([]byte(event), &labelEvent)
	if err != nil {
		t.Error(err)
	} else {
		if labelEvent.Changes["name"].From != "foobar" {
			t.Fatal("incorrect chanages found")
		}
	}
}
