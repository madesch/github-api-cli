package types

//UserInfo encapsulates user information
type UserInfo struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	PublicRepos int    `json:"public_repos"`
}

//Follower encapsulates follower meta
type follower struct {
	Name    string `json:"login"`
	HTMLURL string `json:"html_url"`
}

//Followers represents list of followers
type Followers []follower

//followingUser contains following user information
type followingUser struct {
	Name string `json:"login"`
	URL  string `json:"html_url"`
}

//FollowingUsers stores a list of FollowingUser
type FollowingUsers = []followingUser
