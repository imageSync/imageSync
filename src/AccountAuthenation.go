package src

import "github.com/spf13/viper"

type User struct {
	Username      string
	Password      string
	ServerAddress string
	ImageTag      string
}

type Option func(u *User)

func WithUsername(username string) Option {
	return func(u *User) {
		u.Username = viper.GetString(username)
	}
}

func WithPassword(password string) Option {
	return func(u *User) {
		u.Password = viper.GetString(password)
	}
}

func WithServerAddress(serverAddress string) Option {
	return func(u *User) {
		u.ServerAddress = viper.GetString(serverAddress)
	}
}

func WithImageTag(imageTag string) Option {
	return func(u *User) {
		u.ImageTag = viper.GetString(imageTag)
	}
}

func NewUser(opts ...Option) User {
	user := User{}
	for _, opt := range opts {
		opt(&user)
	}
	return user
}
