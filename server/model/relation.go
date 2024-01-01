package model

// ModeratorsSubforums represents the many-to-many relationship join table
type ModeratorsSubforums struct {
	UserID     uint `json:"UserID"`
	SubforumID uint `json:"SubforumID"`
}

type ModeratorsSubforumsList struct {
	ModeratorsSubforumsList []ModeratorsSubforums `json:"ModeratorsSubforums"`
}

// PostsTags represents the many-to-many relationship join table
type PostsTags struct {
	PostID     uint `json:"PostID"`
	TagID uint `json:"TagID"`
}

type PostsTagsList struct {
	ModeratorsSubforumsList []PostsTags `json:"PostsTags"`
}