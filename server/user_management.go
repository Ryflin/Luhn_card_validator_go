package main

type luser struct {
	Username string `json:"username"`
	Other_info string `json:"other_info"`
	Bio string `json:"bio"`
}
type luserOptions struct {
	other_info string
	bio string
}	
type newLuserOptions func(*luserOptions)

func CreateNewLuser(username string, options ...newLuserOptions) luser {
	opts := luserOptions{}
	for _, option := range options {
		option(&opts)
	}
	return luser{
		Username: username,
		Other_info: opts.other_info,
		Bio: opts.bio,
	}
}