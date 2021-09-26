package model


type ProfilePost struct {
	Profile Profiles
	Posts    []Posts
}

type PostCategory struct {
	Post Posts
	Category Categories
}
