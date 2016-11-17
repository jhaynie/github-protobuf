package github

import (
	"testing"
	"encoding/json"
)

func TestCommitEvent(t *testing.T) {
	event := `{
		"author": {
			"avatar_url": "https://avatars.githubusercontent.com/u/534930?v=3",
			"bio": null,
			"blog": null,
			"company": "Google",
			"created_at": "2010-12-23T19:49:07Z",
			"email": null,
			"followers": 370,
			"following": 9,
			"hireable": null,
			"id": 534930,
			"location": "Mountain View CA",
			"login": "robwormald",
			"name": "Rob Wormald",
			"public_gists": 575,
			"public_repos": 151,
			"site_admin": false,
			"type": "User",
			"updated_at": "2016-07-14T22:20:00Z"
		},
		"commit": {
			"author": {
				"date": "2016-07-01T03:24:39Z",
				"email": "robwormald@xxx.com",
				"name": "Rob Wormald"
			},
			"comment_count": 0,
			"committer": {
				"date": "2016-07-01T03:24:39Z",
				"email": "noreply@github.com",
				"name": "GitHub"
			},
			"message": "fix(forms): use change event for select multiple (#9713)",
			"tree": {
				"sha": "a1168058173b141408bd4034ce7e4675c1c9d992",
				"url": "https://api.github.com/repos/angular/angular/git/trees/a1168058173b141408bd4034ce7e4675c1c9d992"
			},
			"url": "https://api.github.com/repos/angular/angular/git/commits/3cbded6694880dcdf4394ee27f8ab5ca54d9acfd"
		},
		"committer": {
			"avatar_url": "https://avatars.githubusercontent.com/u/534930?v=3",
			"bio": null,
			"blog": null,
			"company": "Google",
			"created_at": "2010-12-23T19:49:07Z",
			"email": null,
			"followers": 370,
			"following": 9,
			"hireable": null,
			"id": 534930,
			"location": "Mountain View CA",
			"login": "robwormald",
			"name": "Rob Wormald",
			"public_gists": 575,
			"public_repos": 151,
			"site_admin": false,
			"type": "User",
			"updated_at": "2016-07-14T22:20:00Z"
		},
		"files": [
			{
				"additions": 2,
				"author": {
					"avatar_url": "https://avatars.githubusercontent.com/u/534930?v=3",
					"bio": null,
					"blog": null,
					"company": "Google",
					"created_at": "2010-12-23T19:49:07Z",
					"email": null,
					"followers": 370,
					"following": 9,
					"hireable": null,
					"id": 534930,
					"location": "Mountain View CA",
					"login": "robwormald",
					"name": "Rob Wormald",
					"public_gists": 575,
					"public_repos": 151,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-07-14T22:20:00Z"
				},
				"binary": false,
				"body": "",
				"commit": "3cbded6694880dcdf4394ee27f8ab5ca54d9acfd",
				"committer": {
					"avatar_url": "https://avatars.githubusercontent.com/u/534930?v=3",
					"bio": null,
					"blog": null,
					"company": "Google",
					"created_at": "2010-12-23T19:49:07Z",
					"email": null,
					"followers": 370,
					"following": 9,
					"hireable": null,
					"id": 534930,
					"location": "Mountain View CA",
					"login": "robwormald",
					"name": "Rob Wormald",
					"public_gists": 575,
					"public_repos": 151,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-07-14T22:20:00Z"
				},
				"deletions": 2,
				"filename": "modules/@angular/common/src/forms-deprecated/directives/select_multiple_control_value_accessor.ts",
				"frameworks": [],
				"id": "713574859313de2f29df9b0ae04886a25220e3ed",
				"language": "TypeScript",
				"licenses": {},
				"linguist": {
					"content_type": "video/MP2T",
					"disposition": "inline",
					"extname": ".ts",
					"is_binary": false,
					"is_documentation": false,
					"is_generated": false,
					"is_high_ratio_of_long_lines": true,
					"is_image": false,
					"is_large": false,
					"is_safe_to_colorize": false,
					"is_text": true,
					"is_vendored": false,
					"is_viewable": true,
					"language": {
						"ace_mode": "typescript",
						"group": "TypeScript",
						"is_popular": false,
						"is_unpopular": true,
						"type": "programming"
					},
					"loc": 1,
					"sloc": 1,
					"type": "text"
				},
				"mimetype": "video/MP2T",
				"patch": "",
				"size": 5437,
				"status": "modified"
			},
			{
				"additions": 2,
				"author": {
					"avatar_url": "https://avatars.githubusercontent.com/u/534930?v=3",
					"bio": null,
					"blog": null,
					"company": "Google",
					"created_at": "2010-12-23T19:49:07Z",
					"email": null,
					"followers": 370,
					"following": 9,
					"hireable": null,
					"id": 534930,
					"location": "Mountain View CA",
					"login": "robwormald",
					"name": "Rob Wormald",
					"public_gists": 575,
					"public_repos": 151,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-07-14T22:20:00Z"
				},
				"binary": false,
				"body": "",
				"commit": "3cbded6694880dcdf4394ee27f8ab5ca54d9acfd",
				"committer": {
					"avatar_url": "https://avatars.githubusercontent.com/u/534930?v=3",
					"bio": null,
					"blog": null,
					"company": "Google",
					"created_at": "2010-12-23T19:49:07Z",
					"email": null,
					"followers": 370,
					"following": 9,
					"hireable": null,
					"id": 534930,
					"location": "Mountain View CA",
					"login": "robwormald",
					"name": "Rob Wormald",
					"public_gists": 575,
					"public_repos": 151,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-07-14T22:20:00Z"
				},
				"deletions": 2,
				"filename": "modules/@angular/forms/src/directives/select_multiple_control_value_accessor.ts",
				"frameworks": [],
				"id": "40dc7b21109b4a30bfa27f0fa5539178dc44760d",
				"language": "TypeScript",
				"licenses": {},
				"linguist": {
					"content_type": "video/MP2T",
					"disposition": "inline",
					"extname": ".ts",
					"is_binary": false,
					"is_documentation": false,
					"is_generated": false,
					"is_high_ratio_of_long_lines": true,
					"is_image": false,
					"is_large": false,
					"is_safe_to_colorize": false,
					"is_text": true,
					"is_vendored": false,
					"is_viewable": true,
					"language": {
						"ace_mode": "typescript",
						"group": "TypeScript",
						"is_popular": false,
						"is_unpopular": true,
						"type": "programming"
					},
					"loc": 1,
					"sloc": 1,
					"type": "text"
				},
				"mimetype": "video/MP2T",
				"patch": "",
				"size": 5464,
				"status": "modified"
			}
		],
		"id": "3cbded6694880dcdf4394ee27f8ab5ca54d9acfd",
		"parents": [
			{
				"sha": "137fff963284de0fe3a20d1f4b2d3a0742db8edc",
				"url": "https://api.github.com/repos/angular/angular/commits/137fff963284de0fe3a20d1f4b2d3a0742db8edc"
			}
		],
		"repo": {
			"archive_url": "https://api.github.com/repos/angular/angular/{archive_format}{/ref}",
			"assignees_url": "https://api.github.com/repos/angular/angular/assignees{/user}",
			"blobs_url": "https://api.github.com/repos/angular/angular/git/blobs{/sha}",
			"branches_url": "https://api.github.com/repos/angular/angular/branches{/branch}",
			"clone_url": "https://github.com/angular/angular.git",
			"collaborators_url": "https://api.github.com/repos/angular/angular/collaborators{/collaborator}",
			"comments_url": "https://api.github.com/repos/angular/angular/comments{/number}",
			"commits_url": "https://api.github.com/repos/angular/angular/commits{/sha}",
			"compare_url": "https://api.github.com/repos/angular/angular/compare/{base}...{head}",
			"contents_url": "https://api.github.com/repos/angular/angular/contents/{+path}",
			"contributors_url": "https://api.github.com/repos/angular/angular/contributors",
			"created_at": "2014-09-18T16:12:01Z",
			"default_branch": "master",
			"deployments_url": "https://api.github.com/repos/angular/angular/deployments",
			"description": "",
			"downloads_url": "https://api.github.com/repos/angular/angular/downloads",
			"events_url": "https://api.github.com/repos/angular/angular/events",
			"fork": false,
			"forks": 4295,
			"forks_count": 4295,
			"forks_url": "https://api.github.com/repos/angular/angular/forks",
			"full_name": "angular/angular",
			"git_commits_url": "https://api.github.com/repos/angular/angular/git/commits{/sha}",
			"git_refs_url": "https://api.github.com/repos/angular/angular/git/refs{/sha}",
			"git_tags_url": "https://api.github.com/repos/angular/angular/git/tags{/sha}",
			"git_url": "git://github.com/angular/angular.git",
			"has_downloads": true,
			"has_issues": true,
			"has_pages": false,
			"has_wiki": false,
			"homepage": "https://angular.io",
			"hooks_url": "https://api.github.com/repos/angular/angular/hooks",
			"html_url": "https://github.com/angular/angular",
			"id": 24195339,
			"issue_comment_url": "https://api.github.com/repos/angular/angular/issues/comments{/number}",
			"issue_events_url": "https://api.github.com/repos/angular/angular/issues/events{/number}",
			"issues_url": "https://api.github.com/repos/angular/angular/issues{/number}",
			"keys_url": "https://api.github.com/repos/angular/angular/keys{/key_id}",
			"labels_url": "https://api.github.com/repos/angular/angular/labels{/name}",
			"language": "TypeScript",
			"languages_url": "https://api.github.com/repos/angular/angular/languages",
			"merges_url": "https://api.github.com/repos/angular/angular/merges",
			"milestones_url": "https://api.github.com/repos/angular/angular/milestones{/number}",
			"mirror_url": null,
			"name": "angular",
			"network_count": 4295,
			"notifications_url": "https://api.github.com/repos/angular/angular/notifications{?since,all,participating}",
			"open_issues": 1029,
			"open_issues_count": 1029,
			"organization": {
				"avatar_url": "https://avatars.githubusercontent.com/u/139426?v=3",
				"events_url": "https://api.github.com/users/angular/events{/privacy}",
				"followers_url": "https://api.github.com/users/angular/followers",
				"following_url": "https://api.github.com/users/angular/following{/other_user}",
				"gists_url": "https://api.github.com/users/angular/gists{/gist_id}",
				"gravatar_id": "",
				"html_url": "https://github.com/angular",
				"id": 139426,
				"login": "angular",
				"organizations_url": "https://api.github.com/users/angular/orgs",
				"received_events_url": "https://api.github.com/users/angular/received_events",
				"repos_url": "https://api.github.com/users/angular/repos",
				"site_admin": false,
				"starred_url": "https://api.github.com/users/angular/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/angular/subscriptions",
				"type": "Organization",
				"url": "https://api.github.com/users/angular"
			},
			"owner": {
				"avatar_url": "https://avatars.githubusercontent.com/u/139426?v=3",
				"events_url": "https://api.github.com/users/angular/events{/privacy}",
				"followers_url": "https://api.github.com/users/angular/followers",
				"following_url": "https://api.github.com/users/angular/following{/other_user}",
				"gists_url": "https://api.github.com/users/angular/gists{/gist_id}",
				"gravatar_id": "",
				"html_url": "https://github.com/angular",
				"id": 139426,
				"login": "angular",
				"organizations_url": "https://api.github.com/users/angular/orgs",
				"received_events_url": "https://api.github.com/users/angular/received_events",
				"repos_url": "https://api.github.com/users/angular/repos",
				"site_admin": false,
				"starred_url": "https://api.github.com/users/angular/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/angular/subscriptions",
				"type": "Organization",
				"url": "https://api.github.com/users/angular"
			},
			"permissions": {
				"admin": false,
				"pull": true,
				"push": false
			},
			"private": false,
			"pulls_url": "https://api.github.com/repos/angular/angular/pulls{/number}",
			"pushed_at": "2016-10-04T02:44:12Z",
			"releases_url": "https://api.github.com/repos/angular/angular/releases{/id}",
			"size": 646139,
			"ssh_url": "git@github.com:angular/angular.git",
			"stargazers_count": 16964,
			"stargazers_url": "https://api.github.com/repos/angular/angular/stargazers",
			"statuses_url": "https://api.github.com/repos/angular/angular/statuses/{sha}",
			"subscribers_count": 2059,
			"subscribers_url": "https://api.github.com/repos/angular/angular/subscribers",
			"subscription_url": "https://api.github.com/repos/angular/angular/subscription",
			"svn_url": "https://github.com/angular/angular",
			"tags_url": "https://api.github.com/repos/angular/angular/tags",
			"teams_url": "https://api.github.com/repos/angular/angular/teams",
			"trees_url": "https://api.github.com/repos/angular/angular/git/trees{/sha}",
			"updated_at": "2016-10-04T06:46:35Z",
			"url": "https://api.github.com/repos/angular/angular",
			"watchers": 16964,
			"watchers_count": 16964
		},
		"sha": "3cbded6694880dcdf4394ee27f8ab5ca54d9acfd"
	}`
	var commit Commit
	err := json.Unmarshal([]byte(event), &commit)
	if err != nil {
		t.Error(err)
	} else {
		if commit.Id != "3cbded6694880dcdf4394ee27f8ab5ca54d9acfd" {
			t.Error("commit.Id != 3cbded6694880dcdf4394ee27f8ab5ca54d9acfd")
		}
		if commit.Sha != "3cbded6694880dcdf4394ee27f8ab5ca54d9acfd" {
			t.Error("commit.Sha != 3cbded6694880dcdf4394ee27f8ab5ca54d9acfd")
		}
		if commit.Author.Company != "Google" {
			t.Error("commit.Author.Company != Google")
		}
		if commit.Commit.Author.Email != "robwormald@xxx.com" {
			t.Error("commit.Commit.Author.Email != robwormald@xxx.com")
		}
		if commit.Repo.Id != 24195339 {
			t.Error("commit.Repo.Id != 24195339")
		}
	}
}

func TestCommitEventWithFrameworksAndLicenses(t *testing.T) {
	event := `{
		"author": {
			"avatar_url": "https://avatars.githubusercontent.com/u/474988?v=3",
			"bio": null,
			"blog": "briantford.com",
			"company": null,
			"created_at": "2010-11-10T01:14:25Z",
			"email": null,
			"followers": 2964,
			"following": 52,
			"hireable": null,
			"id": 474988,
			"location": "San Francisco, CA",
			"login": "btford",
			"name": "Brian Ford",
			"public_gists": 5,
			"public_repos": 249,
			"site_admin": false,
			"type": "User",
			"updated_at": "2016-08-26T03:42:45Z"
		},
		"commit": {
			"author": {
				"date": "2016-01-29T21:28:59Z",
				"email": "btford@xxx.com",
				"name": "Brian Ford"
			},
			"comment_count": 0,
			"committer": {
				"date": "2016-02-02T21:58:51Z",
				"email": "alexeagle@xxx.com",
				"name": "Alex Eagle"
			},
			"message": "build(router): refactor angular1 router build script",
			"tree": {
				"sha": "1a8bc0ee21588017d49c3f28ac893dca382d85bc",
				"url": "https://api.github.com/repos/angular/angular/git/trees/1a8bc0ee21588017d49c3f28ac893dca382d85bc"
			},
			"url": "https://api.github.com/repos/angular/angular/git/commits/6acc99729c7b3f1ffa480fb7e5b585754f197e17"
		},
		"committer": {
			"avatar_url": "https://avatars.githubusercontent.com/u/47395?v=3",
			"bio": null,
			"blog": "http://angular.io",
			"company": "Angular (@Google)",
			"created_at": "2009-01-18T00:40:37Z",
			"email": "eagle@xxx.com",
			"followers": 85,
			"following": 5,
			"hireable": null,
			"id": 47395,
			"location": "SF Bay Area",
			"login": "alexeagle",
			"name": "Alex Eagle",
			"public_gists": 6,
			"public_repos": 37,
			"site_admin": false,
			"type": "User",
			"updated_at": "2016-05-10T22:31:18Z"
		},
		"files": [
			{
				"additions": 5,
				"author": {
					"avatar_url": "https://avatars.githubusercontent.com/u/474988?v=3",
					"bio": null,
					"blog": "briantford.com",
					"company": null,
					"created_at": "2010-11-10T01:14:25Z",
					"email": null,
					"followers": 2964,
					"following": 52,
					"hireable": null,
					"id": 474988,
					"location": "San Francisco, CA",
					"login": "btford",
					"name": "Brian Ford",
					"public_gists": 5,
					"public_repos": 249,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-08-26T03:42:45Z"
				},
				"binary": false,
				"body": "",
				"commit": "6acc99729c7b3f1ffa480fb7e5b585754f197e17",
				"committer": {
					"avatar_url": "https://avatars.githubusercontent.com/u/47395?v=3",
					"bio": null,
					"blog": "http://angular.io",
					"company": "Angular (@Google)",
					"created_at": "2009-01-18T00:40:37Z",
					"email": "eagle@xxx.com",
					"followers": 85,
					"following": 5,
					"hireable": null,
					"id": 47395,
					"location": "SF Bay Area",
					"login": "alexeagle",
					"name": "Alex Eagle",
					"public_gists": 6,
					"public_repos": 37,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-05-10T22:31:18Z"
				},
				"deletions": 1,
				"filename": "gulpfile.js",
				"frameworks": [
					{
						"name": "gulp"
					},
					{
						"name": "uglify"
					}
				],
				"id": "1e3c6de470e54bfebf1a74fa1bbfd2683a9e7d48",
				"language": "JavaScript",
				"licenses": {
					"MIT": {
						"name": "MIT License",
						"url": "http://www.opensource.org/licenses/MIT",
						"osiApproved": true,
						"id": "MIT"
				   }
				},
				"linguist": {
					"content_type": "text/plain; charset=iso-8859-1",
					"disposition": "inline",
					"extname": ".js",
					"is_binary": false,
					"is_documentation": false,
					"is_generated": true,
					"is_high_ratio_of_long_lines": true,
					"is_image": false,
					"is_large": false,
					"is_safe_to_colorize": false,
					"is_text": true,
					"is_vendored": false,
					"is_viewable": true,
					"language": {
						"ace_mode": "javascript",
						"group": "JavaScript",
						"is_popular": true,
						"is_unpopular": false,
						"type": "programming"
					},
					"loc": 1,
					"sloc": 1,
					"type": "text"
				},
				"mimetype": "text/javascript",
				"patch": "",
				"size": 55746,
				"status": "modified"
			},
			{
				"additions": 15,
				"author": {
					"avatar_url": "https://avatars.githubusercontent.com/u/474988?v=3",
					"bio": null,
					"blog": "briantford.com",
					"company": null,
					"created_at": "2010-11-10T01:14:25Z",
					"email": null,
					"followers": 2964,
					"following": 52,
					"hireable": null,
					"id": 474988,
					"location": "San Francisco, CA",
					"login": "btford",
					"name": "Brian Ford",
					"public_gists": 5,
					"public_repos": 249,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-08-26T03:42:45Z"
				},
				"binary": false,
				"body": "",
				"commit": "6acc99729c7b3f1ffa480fb7e5b585754f197e17",
				"committer": {
					"avatar_url": "https://avatars.githubusercontent.com/u/47395?v=3",
					"bio": null,
					"blog": "http://angular.io",
					"company": "Angular (@Google)",
					"created_at": "2009-01-18T00:40:37Z",
					"email": "eagle@xxx.com",
					"followers": 85,
					"following": 5,
					"hireable": null,
					"id": 47395,
					"location": "SF Bay Area",
					"login": "alexeagle",
					"name": "Alex Eagle",
					"public_gists": 6,
					"public_repos": 37,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-05-10T22:31:18Z"
				},
				"deletions": 12,
				"filename": "modules/angular1_router/build.js",
				"frameworks": [
					{
						"name": "uglify"
					}
				],
				"id": "5d67dab6df0a8b4aea3724ea59b524d83c2cd703",
				"language": "JavaScript",
				"licenses": {},
				"linguist": {
					"content_type": "text/plain; charset=iso-8859-1",
					"disposition": "inline",
					"extname": ".js",
					"is_binary": false,
					"is_documentation": false,
					"is_generated": true,
					"is_high_ratio_of_long_lines": false,
					"is_image": false,
					"is_large": false,
					"is_safe_to_colorize": true,
					"is_text": true,
					"is_vendored": true,
					"is_viewable": true,
					"language": {
						"ace_mode": "javascript",
						"group": "JavaScript",
						"is_popular": true,
						"is_unpopular": false,
						"type": "programming"
					},
					"loc": 1,
					"sloc": 1,
					"type": "text"
				},
				"mimetype": "text/javascript",
				"patch": "",
				"size": 2468,
				"status": "modified"
			}
		],
		"id": "6acc99729c7b3f1ffa480fb7e5b585754f197e17",
		"parents": [
			{
				"sha": "99e6500a2d862729654e91db5a9f24ead306db1a",
				"url": "https://api.github.com/repos/angular/angular/commits/99e6500a2d862729654e91db5a9f24ead306db1a"
			}
		],
		"repo": {
			"archive_url": "https://api.github.com/repos/angular/angular/{archive_format}{/ref}",
			"assignees_url": "https://api.github.com/repos/angular/angular/assignees{/user}",
			"blobs_url": "https://api.github.com/repos/angular/angular/git/blobs{/sha}",
			"branches_url": "https://api.github.com/repos/angular/angular/branches{/branch}",
			"clone_url": "https://github.com/angular/angular.git",
			"collaborators_url": "https://api.github.com/repos/angular/angular/collaborators{/collaborator}",
			"comments_url": "https://api.github.com/repos/angular/angular/comments{/number}",
			"commits_url": "https://api.github.com/repos/angular/angular/commits{/sha}",
			"compare_url": "https://api.github.com/repos/angular/angular/compare/{base}...{head}",
			"contents_url": "https://api.github.com/repos/angular/angular/contents/{+path}",
			"contributors_url": "https://api.github.com/repos/angular/angular/contributors",
			"created_at": "2014-09-18T16:12:01Z",
			"default_branch": "master",
			"deployments_url": "https://api.github.com/repos/angular/angular/deployments",
			"description": "",
			"downloads_url": "https://api.github.com/repos/angular/angular/downloads",
			"events_url": "https://api.github.com/repos/angular/angular/events",
			"fork": false,
			"forks": 4296,
			"forks_count": 4296,
			"forks_url": "https://api.github.com/repos/angular/angular/forks",
			"full_name": "angular/angular",
			"git_commits_url": "https://api.github.com/repos/angular/angular/git/commits{/sha}",
			"git_refs_url": "https://api.github.com/repos/angular/angular/git/refs{/sha}",
			"git_tags_url": "https://api.github.com/repos/angular/angular/git/tags{/sha}",
			"git_url": "git://github.com/angular/angular.git",
			"has_downloads": true,
			"has_issues": true,
			"has_pages": false,
			"has_wiki": false,
			"homepage": "https://angular.io",
			"hooks_url": "https://api.github.com/repos/angular/angular/hooks",
			"html_url": "https://github.com/angular/angular",
			"id": 24195339,
			"issue_comment_url": "https://api.github.com/repos/angular/angular/issues/comments{/number}",
			"issue_events_url": "https://api.github.com/repos/angular/angular/issues/events{/number}",
			"issues_url": "https://api.github.com/repos/angular/angular/issues{/number}",
			"keys_url": "https://api.github.com/repos/angular/angular/keys{/key_id}",
			"labels_url": "https://api.github.com/repos/angular/angular/labels{/name}",
			"language": "TypeScript",
			"languages_url": "https://api.github.com/repos/angular/angular/languages",
			"merges_url": "https://api.github.com/repos/angular/angular/merges",
			"milestones_url": "https://api.github.com/repos/angular/angular/milestones{/number}",
			"mirror_url": null,
			"name": "angular",
			"network_count": 4296,
			"notifications_url": "https://api.github.com/repos/angular/angular/notifications{?since,all,participating}",
			"open_issues": 1029,
			"open_issues_count": 1029,
			"organization": {
				"avatar_url": "https://avatars.githubusercontent.com/u/139426?v=3",
				"events_url": "https://api.github.com/users/angular/events{/privacy}",
				"followers_url": "https://api.github.com/users/angular/followers",
				"following_url": "https://api.github.com/users/angular/following{/other_user}",
				"gists_url": "https://api.github.com/users/angular/gists{/gist_id}",
				"gravatar_id": "",
				"html_url": "https://github.com/angular",
				"id": 139426,
				"login": "angular",
				"organizations_url": "https://api.github.com/users/angular/orgs",
				"received_events_url": "https://api.github.com/users/angular/received_events",
				"repos_url": "https://api.github.com/users/angular/repos",
				"site_admin": false,
				"starred_url": "https://api.github.com/users/angular/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/angular/subscriptions",
				"type": "Organization",
				"url": "https://api.github.com/users/angular"
			},
			"owner": {
				"avatar_url": "https://avatars.githubusercontent.com/u/139426?v=3",
				"events_url": "https://api.github.com/users/angular/events{/privacy}",
				"followers_url": "https://api.github.com/users/angular/followers",
				"following_url": "https://api.github.com/users/angular/following{/other_user}",
				"gists_url": "https://api.github.com/users/angular/gists{/gist_id}",
				"gravatar_id": "",
				"html_url": "https://github.com/angular",
				"id": 139426,
				"login": "angular",
				"organizations_url": "https://api.github.com/users/angular/orgs",
				"received_events_url": "https://api.github.com/users/angular/received_events",
				"repos_url": "https://api.github.com/users/angular/repos",
				"site_admin": false,
				"starred_url": "https://api.github.com/users/angular/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/angular/subscriptions",
				"type": "Organization",
				"url": "https://api.github.com/users/angular"
			},
			"permissions": {
				"admin": false,
				"pull": true,
				"push": false
			},
			"private": false,
			"pulls_url": "https://api.github.com/repos/angular/angular/pulls{/number}",
			"pushed_at": "2016-10-04T02:44:12Z",
			"releases_url": "https://api.github.com/repos/angular/angular/releases{/id}",
			"size": 646139,
			"ssh_url": "git@github.com:angular/angular.git",
			"stargazers_count": 16964,
			"stargazers_url": "https://api.github.com/repos/angular/angular/stargazers",
			"statuses_url": "https://api.github.com/repos/angular/angular/statuses/{sha}",
			"subscribers_count": 2059,
			"subscribers_url": "https://api.github.com/repos/angular/angular/subscribers",
			"subscription_url": "https://api.github.com/repos/angular/angular/subscription",
			"svn_url": "https://github.com/angular/angular",
			"tags_url": "https://api.github.com/repos/angular/angular/tags",
			"teams_url": "https://api.github.com/repos/angular/angular/teams",
			"trees_url": "https://api.github.com/repos/angular/angular/git/trees{/sha}",
			"updated_at": "2016-10-04T06:46:35Z",
			"url": "https://api.github.com/repos/angular/angular",
			"watchers": 16964,
			"watchers_count": 16964
		},
		"sha": "6acc99729c7b3f1ffa480fb7e5b585754f197e17"
	}`
	var commit Commit
	err := json.Unmarshal([]byte(event), &commit)
	if err != nil {
		t.Error(err)
	} else {
		if commit.Id != "6acc99729c7b3f1ffa480fb7e5b585754f197e17" {
			t.Error("commit.Id != 6acc99729c7b3f1ffa480fb7e5b585754f197e17")
		}
		if len(commit.Files[0].Frameworks) != 2 {
			t.Error("len(commit.Files[0].Frameworks) != 2")
		}
		if commit.Files[0].Frameworks[0].Name != "gulp" {
			t.Error("commit.Files[0].Frameworks[0].Name != gulp")
		}
		if len(commit.Files[0].Licenses) != 1 {
			t.Error("len(commit.Files[0].Licenses) != 1")
		}
		lic := commit.Files[0].Licenses["MIT"]
		if lic == nil {
			t.Error("commit.Files[0].Licenses[\"MIT\"] == nil")
		}
		if lic.Id != "MIT" {
			t.Error("commit.Files[0].Licenses[\"MIT\"].Id != MIT")
		}
		if lic.Name != "MIT License" {
			t.Error("commit.Files[0].Licenses[\"MIT\"].Name != MIT License")
		}
		if lic.OsiApproved == false {
			t.Error("commit.Files[0].Licenses[\"MIT\"].OsiApproved != true")
		}
		if lic.Url != "http://www.opensource.org/licenses/MIT" {
			t.Error("commit.Files[0].Licenses[\"MIT\"].Url != http://www.opensource.org/licenses/MIT")
		}
		if commit.Commit.Author.Date != "2016-01-29T21:28:59Z" {
			t.Error("commit.Commit.Author.Date != 2016-01-29T21:28:59Z")
		}
		if commit.Files[0].FileSize != 55746 {
			t.Error("commit.Files[0].FileSize != 55746")
		}
		if commit.Files[1].FileSize != 2468 {
			t.Error("commit.Files[1].FileSize != 2468")
		}
	}
}

func TestCommitEventWithProprietaryLicense(t *testing.T) {
	event := `{
		"author": {
			"avatar_url": "https://avatars.githubusercontent.com/u/474988?v=3",
			"bio": null,
			"blog": "briantford.com",
			"company": null,
			"created_at": "2010-11-10T01:14:25Z",
			"email": null,
			"followers": 2964,
			"following": 52,
			"hireable": null,
			"id": 474988,
			"location": "San Francisco, CA",
			"login": "btford",
			"name": "Brian Ford",
			"public_gists": 5,
			"public_repos": 249,
			"site_admin": false,
			"type": "User",
			"updated_at": "2016-08-26T03:42:45Z"
		},
		"commit": {
			"author": {
				"date": "2016-01-29T21:28:59Z",
				"email": "btford@xxx.com",
				"name": "Brian Ford"
			},
			"comment_count": 0,
			"committer": {
				"date": "2016-02-02T21:58:51Z",
				"email": "alexeagle@xxx.com",
				"name": "Alex Eagle"
			},
			"message": "build(router): refactor angular1 router build script",
			"tree": {
				"sha": "1a8bc0ee21588017d49c3f28ac893dca382d85bc",
				"url": "https://api.github.com/repos/angular/angular/git/trees/1a8bc0ee21588017d49c3f28ac893dca382d85bc"
			},
			"url": "https://api.github.com/repos/angular/angular/git/commits/6acc99729c7b3f1ffa480fb7e5b585754f197e17"
		},
		"committer": {
			"avatar_url": "https://avatars.githubusercontent.com/u/47395?v=3",
			"bio": null,
			"blog": "http://angular.io",
			"company": "Angular (@Google)",
			"created_at": "2009-01-18T00:40:37Z",
			"email": "eagle@xxx.com",
			"followers": 85,
			"following": 5,
			"hireable": null,
			"id": 47395,
			"location": "SF Bay Area",
			"login": "alexeagle",
			"name": "Alex Eagle",
			"public_gists": 6,
			"public_repos": 37,
			"site_admin": false,
			"type": "User",
			"updated_at": "2016-05-10T22:31:18Z"
		},
		"files": [
			{
				"additions": 5,
				"author": {
					"avatar_url": "https://avatars.githubusercontent.com/u/474988?v=3",
					"bio": null,
					"blog": "briantford.com",
					"company": null,
					"created_at": "2010-11-10T01:14:25Z",
					"email": null,
					"followers": 2964,
					"following": 52,
					"hireable": null,
					"id": 474988,
					"location": "San Francisco, CA",
					"login": "btford",
					"name": "Brian Ford",
					"public_gists": 5,
					"public_repos": 249,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-08-26T03:42:45Z"
				},
				"binary": false,
				"body": "",
				"commit": "6acc99729c7b3f1ffa480fb7e5b585754f197e17",
				"committer": {
					"avatar_url": "https://avatars.githubusercontent.com/u/47395?v=3",
					"bio": null,
					"blog": "http://angular.io",
					"company": "Angular (@Google)",
					"created_at": "2009-01-18T00:40:37Z",
					"email": "eagle@xxx.com",
					"followers": 85,
					"following": 5,
					"hireable": null,
					"id": 47395,
					"location": "SF Bay Area",
					"login": "alexeagle",
					"name": "Alex Eagle",
					"public_gists": 6,
					"public_repos": 37,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-05-10T22:31:18Z"
				},
				"deletions": 1,
				"filename": "gulpfile.js",
				"frameworks": [
					{
						"name": "gulp"
					},
					{
						"name": "uglify"
					}
				],
				"id": "1e3c6de470e54bfebf1a74fa1bbfd2683a9e7d48",
				"language": "JavaScript",
				"licenses": {
					"proprietary": {
						"name":"proprietary",
						"osiApproved": false,
						"license": "test"
				   }
				},
				"linguist": {
					"content_type": "text/plain; charset=iso-8859-1",
					"disposition": "inline",
					"extname": ".js",
					"is_binary": false,
					"is_documentation": false,
					"is_generated": true,
					"is_high_ratio_of_long_lines": true,
					"is_image": false,
					"is_large": false,
					"is_safe_to_colorize": false,
					"is_text": true,
					"is_vendored": false,
					"is_viewable": true,
					"language": {
						"ace_mode": "javascript",
						"group": "JavaScript",
						"is_popular": true,
						"is_unpopular": false,
						"type": "programming"
					},
					"loc": 1,
					"sloc": 1,
					"type": "text"
				},
				"mimetype": "text/javascript",
				"patch": "",
				"size": 55746,
				"status": "modified"
			},
			{
				"additions": 15,
				"author": {
					"avatar_url": "https://avatars.githubusercontent.com/u/474988?v=3",
					"bio": null,
					"blog": "briantford.com",
					"company": null,
					"created_at": "2010-11-10T01:14:25Z",
					"email": null,
					"followers": 2964,
					"following": 52,
					"hireable": null,
					"id": 474988,
					"location": "San Francisco, CA",
					"login": "btford",
					"name": "Brian Ford",
					"public_gists": 5,
					"public_repos": 249,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-08-26T03:42:45Z"
				},
				"binary": false,
				"body": "",
				"commit": "6acc99729c7b3f1ffa480fb7e5b585754f197e17",
				"committer": {
					"avatar_url": "https://avatars.githubusercontent.com/u/47395?v=3",
					"bio": null,
					"blog": "http://angular.io",
					"company": "Angular (@Google)",
					"created_at": "2009-01-18T00:40:37Z",
					"email": "eagle@xxx.com",
					"followers": 85,
					"following": 5,
					"hireable": null,
					"id": 47395,
					"location": "SF Bay Area",
					"login": "alexeagle",
					"name": "Alex Eagle",
					"public_gists": 6,
					"public_repos": 37,
					"site_admin": false,
					"type": "User",
					"updated_at": "2016-05-10T22:31:18Z"
				},
				"deletions": 12,
				"filename": "modules/angular1_router/build.js",
				"frameworks": [
					{
						"name": "uglify"
					}
				],
				"id": "5d67dab6df0a8b4aea3724ea59b524d83c2cd703",
				"language": "JavaScript",
				"licenses": {},
				"linguist": {
					"content_type": "text/plain; charset=iso-8859-1",
					"disposition": "inline",
					"extname": ".js",
					"is_binary": false,
					"is_documentation": false,
					"is_generated": true,
					"is_high_ratio_of_long_lines": false,
					"is_image": false,
					"is_large": false,
					"is_safe_to_colorize": true,
					"is_text": true,
					"is_vendored": true,
					"is_viewable": true,
					"language": {
						"ace_mode": "javascript",
						"group": "JavaScript",
						"is_popular": true,
						"is_unpopular": false,
						"type": "programming"
					},
					"loc": 1,
					"sloc": 1,
					"type": "text"
				},
				"mimetype": "text/javascript",
				"patch": "",
				"size": 2468,
				"status": "modified"
			}
		],
		"id": "6acc99729c7b3f1ffa480fb7e5b585754f197e17",
		"parents": [
			{
				"sha": "99e6500a2d862729654e91db5a9f24ead306db1a",
				"url": "https://api.github.com/repos/angular/angular/commits/99e6500a2d862729654e91db5a9f24ead306db1a"
			}
		],
		"repo": {
			"archive_url": "https://api.github.com/repos/angular/angular/{archive_format}{/ref}",
			"assignees_url": "https://api.github.com/repos/angular/angular/assignees{/user}",
			"blobs_url": "https://api.github.com/repos/angular/angular/git/blobs{/sha}",
			"branches_url": "https://api.github.com/repos/angular/angular/branches{/branch}",
			"clone_url": "https://github.com/angular/angular.git",
			"collaborators_url": "https://api.github.com/repos/angular/angular/collaborators{/collaborator}",
			"comments_url": "https://api.github.com/repos/angular/angular/comments{/number}",
			"commits_url": "https://api.github.com/repos/angular/angular/commits{/sha}",
			"compare_url": "https://api.github.com/repos/angular/angular/compare/{base}...{head}",
			"contents_url": "https://api.github.com/repos/angular/angular/contents/{+path}",
			"contributors_url": "https://api.github.com/repos/angular/angular/contributors",
			"created_at": "2014-09-18T16:12:01Z",
			"default_branch": "master",
			"deployments_url": "https://api.github.com/repos/angular/angular/deployments",
			"description": "",
			"downloads_url": "https://api.github.com/repos/angular/angular/downloads",
			"events_url": "https://api.github.com/repos/angular/angular/events",
			"fork": false,
			"forks": 4296,
			"forks_count": 4296,
			"forks_url": "https://api.github.com/repos/angular/angular/forks",
			"full_name": "angular/angular",
			"git_commits_url": "https://api.github.com/repos/angular/angular/git/commits{/sha}",
			"git_refs_url": "https://api.github.com/repos/angular/angular/git/refs{/sha}",
			"git_tags_url": "https://api.github.com/repos/angular/angular/git/tags{/sha}",
			"git_url": "git://github.com/angular/angular.git",
			"has_downloads": true,
			"has_issues": true,
			"has_pages": false,
			"has_wiki": false,
			"homepage": "https://angular.io",
			"hooks_url": "https://api.github.com/repos/angular/angular/hooks",
			"html_url": "https://github.com/angular/angular",
			"id": 24195339,
			"issue_comment_url": "https://api.github.com/repos/angular/angular/issues/comments{/number}",
			"issue_events_url": "https://api.github.com/repos/angular/angular/issues/events{/number}",
			"issues_url": "https://api.github.com/repos/angular/angular/issues{/number}",
			"keys_url": "https://api.github.com/repos/angular/angular/keys{/key_id}",
			"labels_url": "https://api.github.com/repos/angular/angular/labels{/name}",
			"language": "TypeScript",
			"languages_url": "https://api.github.com/repos/angular/angular/languages",
			"merges_url": "https://api.github.com/repos/angular/angular/merges",
			"milestones_url": "https://api.github.com/repos/angular/angular/milestones{/number}",
			"mirror_url": null,
			"name": "angular",
			"network_count": 4296,
			"notifications_url": "https://api.github.com/repos/angular/angular/notifications{?since,all,participating}",
			"open_issues": 1029,
			"open_issues_count": 1029,
			"organization": {
				"avatar_url": "https://avatars.githubusercontent.com/u/139426?v=3",
				"events_url": "https://api.github.com/users/angular/events{/privacy}",
				"followers_url": "https://api.github.com/users/angular/followers",
				"following_url": "https://api.github.com/users/angular/following{/other_user}",
				"gists_url": "https://api.github.com/users/angular/gists{/gist_id}",
				"gravatar_id": "",
				"html_url": "https://github.com/angular",
				"id": 139426,
				"login": "angular",
				"organizations_url": "https://api.github.com/users/angular/orgs",
				"received_events_url": "https://api.github.com/users/angular/received_events",
				"repos_url": "https://api.github.com/users/angular/repos",
				"site_admin": false,
				"starred_url": "https://api.github.com/users/angular/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/angular/subscriptions",
				"type": "Organization",
				"url": "https://api.github.com/users/angular"
			},
			"owner": {
				"avatar_url": "https://avatars.githubusercontent.com/u/139426?v=3",
				"events_url": "https://api.github.com/users/angular/events{/privacy}",
				"followers_url": "https://api.github.com/users/angular/followers",
				"following_url": "https://api.github.com/users/angular/following{/other_user}",
				"gists_url": "https://api.github.com/users/angular/gists{/gist_id}",
				"gravatar_id": "",
				"html_url": "https://github.com/angular",
				"id": 139426,
				"login": "angular",
				"organizations_url": "https://api.github.com/users/angular/orgs",
				"received_events_url": "https://api.github.com/users/angular/received_events",
				"repos_url": "https://api.github.com/users/angular/repos",
				"site_admin": false,
				"starred_url": "https://api.github.com/users/angular/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/angular/subscriptions",
				"type": "Organization",
				"url": "https://api.github.com/users/angular"
			},
			"permissions": {
				"admin": false,
				"pull": true,
				"push": false
			},
			"private": false,
			"pulls_url": "https://api.github.com/repos/angular/angular/pulls{/number}",
			"pushed_at": "2016-10-04T02:44:12Z",
			"releases_url": "https://api.github.com/repos/angular/angular/releases{/id}",
			"size": 646139,
			"ssh_url": "git@github.com:angular/angular.git",
			"stargazers_count": 16964,
			"stargazers_url": "https://api.github.com/repos/angular/angular/stargazers",
			"statuses_url": "https://api.github.com/repos/angular/angular/statuses/{sha}",
			"subscribers_count": 2059,
			"subscribers_url": "https://api.github.com/repos/angular/angular/subscribers",
			"subscription_url": "https://api.github.com/repos/angular/angular/subscription",
			"svn_url": "https://github.com/angular/angular",
			"tags_url": "https://api.github.com/repos/angular/angular/tags",
			"teams_url": "https://api.github.com/repos/angular/angular/teams",
			"trees_url": "https://api.github.com/repos/angular/angular/git/trees{/sha}",
			"updated_at": "2016-10-04T06:46:35Z",
			"url": "https://api.github.com/repos/angular/angular",
			"watchers": 16964,
			"watchers_count": 16964
		},
		"sha": "6acc99729c7b3f1ffa480fb7e5b585754f197e17"
	}`
	var commit Commit
	err := json.Unmarshal([]byte(event), &commit)
	if err != nil {
		t.Error(err)
	} else {
		if commit.Id != "6acc99729c7b3f1ffa480fb7e5b585754f197e17" {
			t.Error("commit.Id != 6acc99729c7b3f1ffa480fb7e5b585754f197e17")
		}
		if len(commit.Files[0].Frameworks) != 2 {
			t.Error("len(commit.Files[0].Frameworks) != 2")
		}
		if commit.Files[0].Frameworks[0].Name != "gulp" {
			t.Error("commit.Files[0].Frameworks[0].Name != gulp")
		}
		if len(commit.Files[0].Licenses) != 1 {
			t.Error("len(commit.Files[0].Licenses) != 1")
		}
		lic := commit.Files[0].Licenses["proprietary"]
		if lic == nil {
			t.Error("commit.Files[0].Licenses[\"proprietary\"] == nil")
		}
		if lic.Name != "proprietary" {
			t.Error("commit.Files[0].Licenses[\"proprietary\"].Name != proprietary")
		}
		if lic.OsiApproved {
			t.Error("commit.Files[0].Licenses[\"MIT\"].OsiApproved != false")
		}
		if lic.License != "test" {
			t.Error("commit.Files[0].Licenses[\"proprietary\"].License != test")
		}
		if commit.Commit.Author.Date != "2016-01-29T21:28:59Z" {
			t.Error("commit.Commit.Author.Date != 2016-01-29T21:28:59Z")
		}
		if commit.Files[0].FileSize != 55746 {
			t.Error("commit.Files[0].FileSize != 55746")
		}
		if commit.Files[1].FileSize != 2468 {
			t.Error("commit.Files[1].FileSize != 2468")
		}
	}
}
