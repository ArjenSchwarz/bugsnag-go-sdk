package api

import (
	"errors"
	"fmt"
)

type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	AccountAdmin bool   `json:"account_admin"`
	Email        string `json:"email"`
	GravatarId   string `json:"gravatar_id"`
	GravatarUrl  string `json:"gravatar_url"`
	HtmlUrl      string `json:"html_url"`
	Url          string `json:"url"`
}

func (c *Connection) User() (User, error) {
	var resp User
	if c.client.Credentials.Username == "" {
		return resp, errors.New("This method can only be used when authenticated with username and password")
	}
	err := c.client.Query("GET", "/user/", &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Connection) Users(object UserList) ([]User, error) {
	var resp []User
	err := c.client.Query("GET", object.UserListUrl(), &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

type UserList interface {
	UserListUrl() string
}

func (c *Connection) UserById(userId string) (User, error) {
	url := fmt.Sprintf("/users/%s", userId)
	var resp User
	err := c.client.Query("GET", url, &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (user *User) ProjectsListUrl() string {
	url := fmt.Sprintf("%s/users/%s/projects", bugsnagApi, user.Id)
	return url
}
