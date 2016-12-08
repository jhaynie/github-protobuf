package github

import (
	"testing"
	"encoding/json"
)

func TestIssuesEvent(t *testing.T) {
	event := `{
	  "action": "opened",
	  "issue": {
	    "url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2",
	    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/labels{/name}",
	    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/comments",
	    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/events",
	    "html_url": "https://github.com/baxterthehacker/public-repo/issues/2",
	    "id": 73464126,
	    "number": 2,
	    "title": "Spelling error in the README file",
	    "user": {
	      "login": "baxterthehacker",
	      "id": 6752317,
	      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/baxterthehacker",
	      "html_url": "https://github.com/baxterthehacker",
	      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
	      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
	      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
	      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
	      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
	      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
	      "type": "User",
	      "site_admin": false
	    },
	    "labels": [
	      {
	        "url": "https://api.github.com/repos/baxterthehacker/public-repo/labels/bug",
	        "name": "bug",
	        "color": "fc2929"
	      }
	    ],
	    "state": "open",
	    "locked": false,
	    "assignee": null,
	    "milestone": null,
	    "comments": 0,
	    "created_at": "2015-05-05T23:40:28Z",
	    "updated_at": "2015-05-05T23:40:28Z",
	    "closed_at": null,
	    "body": "It looks like you accidently spelled 'commit' with two 't's."
	  },
	  "repository": {
	    "id": 35129377,
	    "name": "public-repo",
	    "full_name": "baxterthehacker/public-repo",
	    "owner": {
	      "login": "baxterthehacker",
	      "id": 6752317,
	      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/baxterthehacker",
	      "html_url": "https://github.com/baxterthehacker",
	      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
	      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
	      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
	      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
	      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
	      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
	      "type": "User",
	      "site_admin": false
	    },
	    "private": false,
	    "html_url": "https://github.com/baxterthehacker/public-repo",
	    "description": "",
	    "fork": false,
	    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
	    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
	    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
	    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
	    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
	    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
	    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
	    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
	    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
	    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
	    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
	    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
	    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
	    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
	    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
	    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
	    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
	    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
	    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
	    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
	    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
	    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
	    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
	    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
	    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
	    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
	    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
	    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
	    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
	    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
	    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
	    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
	    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
	    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
	    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
	    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
	    "created_at": "2015-05-05T23:40:12Z",
	    "updated_at": "2015-05-05T23:40:12Z",
	    "pushed_at": "2015-05-05T23:40:27Z",
	    "git_url": "git://github.com/baxterthehacker/public-repo.git",
	    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
	    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
	    "svn_url": "https://github.com/baxterthehacker/public-repo",
	    "homepage": null,
	    "size": 0,
	    "stargazers_count": 0,
	    "watchers_count": 0,
	    "language": null,
	    "has_issues": true,
	    "has_downloads": true,
	    "has_wiki": true,
	    "has_pages": true,
	    "forks_count": 0,
	    "mirror_url": null,
	    "open_issues_count": 2,
	    "forks": 0,
	    "open_issues": 2,
	    "watchers": 0,
	    "default_branch": "master"
	  },
	  "sender": {
	    "login": "baxterthehacker",
	    "id": 6752317,
	    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
	    "gravatar_id": "",
	    "url": "https://api.github.com/users/baxterthehacker",
	    "html_url": "https://github.com/baxterthehacker",
	    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
	    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
	    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
	    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
	    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
	    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
	    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
	    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
	    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
	    "type": "User",
	    "site_admin": false
	  }
	}`
	var issuesEvent IssuesEvent
	err := json.Unmarshal([]byte(event), &issuesEvent)
	if err != nil {
		t.Error(err)
	} else {
		if issuesEvent.Action != "opened" {
			t.Error("issuesEvent.Action != opened")
		}
		if issuesEvent.Issue.Title != "Spelling error in the README file" {
			t.Error("issuesEvent.Issue.Title != Spelling error in the README file")
		}
		if len(issuesEvent.Issue.Labels) != 1 {
			t.Error("len(issuesEvent.Issue.Labels) != 1")
		}
		if issuesEvent.Issue.Labels[0].Color != "fc2929" {
			t.Error("issuesEvent.Issue.Labels[0].Color != fc2929")
		}
		repo := issuesEvent.Repository
		if repo.Id != 35129377 {
			t.Error("issuesEvent.Repository.Id != 35129377")
		}
		if repo.Owner.Login != "baxterthehacker" {
			t.Error("issuesEvent.Repository.Owner.Login != baxterthehacker")
		}
		sender := issuesEvent.Sender
		if sender.Login != "baxterthehacker" {
			t.Error("issuesEvent.Sender.Login != baxterthehacker")
		}
	}
}

func TestIssuesEventWithIntegration(t *testing.T) {
	event := `{
	  "action": "opened",
	  "issue": {
	    "url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/issues/9",
	    "repository_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka",
	    "labels_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/issues/9/labels{/name}",
	    "comments_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/issues/9/comments",
	    "events_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/issues/9/events",
	    "html_url": "https://github.com/jhaynie/pinpoint-webhook-kafka/issues/9",
	    "id": 190896521,
	    "number": 9,
	    "title": "test",
	    "user": {
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
	    },
	    "labels": [

	    ],
	    "state": "open",
	    "locked": false,
	    "assignee": null,
	    "assignees": [

	    ],
	    "milestone": null,
	    "comments": 0,
	    "created_at": "2016-11-22T03:09:53Z",
	    "updated_at": "2016-11-22T03:09:53Z",
	    "closed_at": null,
	    "body": "test"
	  },
	  "repository": {
	    "id": 70075614,
	    "name": "pinpoint-webhook-kafka",
	    "full_name": "jhaynie/pinpoint-webhook-kafka",
	    "owner": {
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
	    },
	    "private": true,
	    "html_url": "https://github.com/jhaynie/pinpoint-webhook-kafka",
	    "description": "new worker stuff",
	    "fork": false,
	    "url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka",
	    "forks_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/forks",
	    "keys_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/keys{/key_id}",
	    "collaborators_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/collaborators{/collaborator}",
	    "teams_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/teams",
	    "hooks_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/hooks",
	    "issue_events_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/issues/events{/number}",
	    "events_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/events",
	    "assignees_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/assignees{/user}",
	    "branches_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/branches{/branch}",
	    "tags_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/tags",
	    "blobs_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/git/blobs{/sha}",
	    "git_tags_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/git/tags{/sha}",
	    "git_refs_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/git/refs{/sha}",
	    "trees_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/git/trees{/sha}",
	    "statuses_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/statuses/{sha}",
	    "languages_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/languages",
	    "stargazers_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/stargazers",
	    "contributors_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/contributors",
	    "subscribers_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/subscribers",
	    "subscription_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/subscription",
	    "commits_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/commits{/sha}",
	    "git_commits_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/git/commits{/sha}",
	    "comments_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/comments{/number}",
	    "issue_comment_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/issues/comments{/number}",
	    "contents_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/contents/{+path}",
	    "compare_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/compare/{base}...{head}",
	    "merges_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/merges",
	    "archive_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/{archive_format}{/ref}",
	    "downloads_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/downloads",
	    "issues_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/issues{/number}",
	    "pulls_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/pulls{/number}",
	    "milestones_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/milestones{/number}",
	    "notifications_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/notifications{?since,all,participating}",
	    "labels_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/labels{/name}",
	    "releases_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/releases{/id}",
	    "deployments_url": "https://api.github.com/repos/jhaynie/pinpoint-webhook-kafka/deployments",
	    "created_at": "2016-10-05T15:54:37Z",
	    "updated_at": "2016-10-24T04:04:11Z",
	    "pushed_at": "2016-11-22T03:01:54Z",
	    "git_url": "git://github.com/jhaynie/pinpoint-webhook-kafka.git",
	    "ssh_url": "git@github.com:jhaynie/pinpoint-webhook-kafka.git",
	    "clone_url": "https://github.com/jhaynie/pinpoint-webhook-kafka.git",
	    "svn_url": "https://github.com/jhaynie/pinpoint-webhook-kafka",
	    "homepage": null,
	    "size": 2880,
	    "stargazers_count": 0,
	    "watchers_count": 0,
	    "language": "Go",
	    "has_issues": true,
	    "has_downloads": true,
	    "has_wiki": true,
	    "has_pages": false,
	    "forks_count": 0,
	    "mirror_url": null,
	    "open_issues_count": 1,
	    "forks": 0,
	    "open_issues": 1,
	    "watchers": 0,
	    "default_branch": "master"
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
	  },
	  "installation": {
	    "id": 3342
	  }
	}`
	var issuesEvent IssuesEvent
	err := json.Unmarshal([]byte(event), &issuesEvent)
	if err != nil {
		t.Error(err)
	} else {
		if issuesEvent.Installation.Id != 3342 {
			t.Error("issuesEvent.Installation.Id != 3342")
		}
	}
}

func TestIssuesEventWithPR(t *testing.T) {
	event := `{
	  "action": "created",
	  "issue": {
	    "url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues/100",
	    "repository_url": "https://api.github.com/repos/jhaynie/pinpoint-web",
	    "labels_url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues/100/labels{/name}",
	    "comments_url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues/100/comments",
	    "events_url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues/100/events",
	    "html_url": "https://github.com/jhaynie/pinpoint-web/pull/100",
	    "id": 194220572,
	    "number": 100,
	    "title": "test for PR",
	    "user": {
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
	    },
	    "labels": [

	    ],
	    "state": "open",
	    "locked": false,
	    "assignee": null,
	    "assignees": [

	    ],
	    "milestone": null,
	    "comments": 0,
	    "created_at": "2016-12-08T01:07:10Z",
	    "updated_at": "2016-12-08T01:07:50Z",
	    "closed_at": null,
	    "pull_request": {
	      "url": "https://api.github.com/repos/jhaynie/pinpoint-web/pulls/100",
	      "html_url": "https://github.com/jhaynie/pinpoint-web/pull/100",
	      "diff_url": "https://github.com/jhaynie/pinpoint-web/pull/100.diff",
	      "patch_url": "https://github.com/jhaynie/pinpoint-web/pull/100.patch"
	    },
	    "body": "test for PR"
	  },
	  "comment": {
	    "url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues/comments/265623454",
	    "html_url": "https://github.com/jhaynie/pinpoint-web/pull/100#issuecomment-265623454",
	    "issue_url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues/100",
	    "id": 265623454,
	    "user": {
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
	    },
	    "created_at": "2016-12-08T01:07:50Z",
	    "updated_at": "2016-12-08T01:07:50Z",
	    "body": "test comment on PR"
	  },
	  "repository": {
	    "id": 63737890,
	    "name": "pinpoint-web",
	    "full_name": "jhaynie/pinpoint-web",
	    "owner": {
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
	    },
	    "private": true,
	    "html_url": "https://github.com/jhaynie/pinpoint-web",
	    "description": "Pinpoint web + api server",
	    "fork": false,
	    "url": "https://api.github.com/repos/jhaynie/pinpoint-web",
	    "forks_url": "https://api.github.com/repos/jhaynie/pinpoint-web/forks",
	    "keys_url": "https://api.github.com/repos/jhaynie/pinpoint-web/keys{/key_id}",
	    "collaborators_url": "https://api.github.com/repos/jhaynie/pinpoint-web/collaborators{/collaborator}",
	    "teams_url": "https://api.github.com/repos/jhaynie/pinpoint-web/teams",
	    "hooks_url": "https://api.github.com/repos/jhaynie/pinpoint-web/hooks",
	    "issue_events_url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues/events{/number}",
	    "events_url": "https://api.github.com/repos/jhaynie/pinpoint-web/events",
	    "assignees_url": "https://api.github.com/repos/jhaynie/pinpoint-web/assignees{/user}",
	    "branches_url": "https://api.github.com/repos/jhaynie/pinpoint-web/branches{/branch}",
	    "tags_url": "https://api.github.com/repos/jhaynie/pinpoint-web/tags",
	    "blobs_url": "https://api.github.com/repos/jhaynie/pinpoint-web/git/blobs{/sha}",
	    "git_tags_url": "https://api.github.com/repos/jhaynie/pinpoint-web/git/tags{/sha}",
	    "git_refs_url": "https://api.github.com/repos/jhaynie/pinpoint-web/git/refs{/sha}",
	    "trees_url": "https://api.github.com/repos/jhaynie/pinpoint-web/git/trees{/sha}",
	    "statuses_url": "https://api.github.com/repos/jhaynie/pinpoint-web/statuses/{sha}",
	    "languages_url": "https://api.github.com/repos/jhaynie/pinpoint-web/languages",
	    "stargazers_url": "https://api.github.com/repos/jhaynie/pinpoint-web/stargazers",
	    "contributors_url": "https://api.github.com/repos/jhaynie/pinpoint-web/contributors",
	    "subscribers_url": "https://api.github.com/repos/jhaynie/pinpoint-web/subscribers",
	    "subscription_url": "https://api.github.com/repos/jhaynie/pinpoint-web/subscription",
	    "commits_url": "https://api.github.com/repos/jhaynie/pinpoint-web/commits{/sha}",
	    "git_commits_url": "https://api.github.com/repos/jhaynie/pinpoint-web/git/commits{/sha}",
	    "comments_url": "https://api.github.com/repos/jhaynie/pinpoint-web/comments{/number}",
	    "issue_comment_url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues/comments{/number}",
	    "contents_url": "https://api.github.com/repos/jhaynie/pinpoint-web/contents/{+path}",
	    "compare_url": "https://api.github.com/repos/jhaynie/pinpoint-web/compare/{base}...{head}",
	    "merges_url": "https://api.github.com/repos/jhaynie/pinpoint-web/merges",
	    "archive_url": "https://api.github.com/repos/jhaynie/pinpoint-web/{archive_format}{/ref}",
	    "downloads_url": "https://api.github.com/repos/jhaynie/pinpoint-web/downloads",
	    "issues_url": "https://api.github.com/repos/jhaynie/pinpoint-web/issues{/number}",
	    "pulls_url": "https://api.github.com/repos/jhaynie/pinpoint-web/pulls{/number}",
	    "milestones_url": "https://api.github.com/repos/jhaynie/pinpoint-web/milestones{/number}",
	    "notifications_url": "https://api.github.com/repos/jhaynie/pinpoint-web/notifications{?since,all,participating}",
	    "labels_url": "https://api.github.com/repos/jhaynie/pinpoint-web/labels{/name}",
	    "releases_url": "https://api.github.com/repos/jhaynie/pinpoint-web/releases{/id}",
	    "deployments_url": "https://api.github.com/repos/jhaynie/pinpoint-web/deployments",
	    "created_at": "2016-07-20T00:50:16Z",
	    "updated_at": "2016-10-24T18:38:43Z",
	    "pushed_at": "2016-12-08T01:07:10Z",
	    "git_url": "git://github.com/jhaynie/pinpoint-web.git",
	    "ssh_url": "git@github.com:jhaynie/pinpoint-web.git",
	    "clone_url": "https://github.com/jhaynie/pinpoint-web.git",
	    "svn_url": "https://github.com/jhaynie/pinpoint-web",
	    "homepage": null,
	    "size": 1225,
	    "stargazers_count": 0,
	    "watchers_count": 0,
	    "language": "JavaScript",
	    "has_issues": true,
	    "has_downloads": true,
	    "has_wiki": false,
	    "has_pages": false,
	    "forks_count": 0,
	    "mirror_url": null,
	    "open_issues_count": 19,
	    "forks": 0,
	    "open_issues": 19,
	    "watchers": 0,
	    "default_branch": "master"
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
	  },
	  "installation": {
	    "id": 4137
	  }
	}`
	var issuesEvent IssuesEvent
	err := json.Unmarshal([]byte(event), &issuesEvent)
	if err != nil {
		t.Error(err)
	} else {
		if issuesEvent.Installation.Id != 4137 {
			t.Error("issuesEvent.Installation.Id != 4137")
		}
		if issuesEvent.Issue.PullRequest.Url != "https://api.github.com/repos/jhaynie/pinpoint-web/pulls/100" {
			t.Error("issuesEvent.Issue.PullRequest.Url != https://api.github.com/repos/jhaynie/pinpoint-web/pulls/100")
		}
	}
}
