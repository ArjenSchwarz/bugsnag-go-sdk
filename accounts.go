package api

import (
	"errors"
	"fmt"
)

type Account struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	AccountCreator User   `json:"account_creator"`
	BillingContact User   `json:"billing_contact"`
	Url            string `json:"url"`
	UsersUrl       string `json:"users_url"`
	ProjectsUrl    string `json:"projects_url"`
	CreatedAt      jTime  `json:"created_at"`
	UpdatedAt      jTime  `json:"updated_at"`
}

func (c *Connection) Accounts() ([]Account, error) {
	var resp []Account
	err := c.client.Query("GET", "/accounts", &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connection) Account() (Account, error) {
	var resp Account
	if c.client.Credentials.ApiKey == "" {
		return resp, errors.New("This method can only be used when authenticated with an API key")
	}
	err := c.client.Query("GET", "/account", &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Connection) AccountById(accountId string) (Account, error) {
	url := fmt.Sprintf("/accounts/%s", accountId)
	var resp Account
	err := c.client.Query("GET", url, &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (account *Account) UserListUrl() string {
	return account.UsersUrl
}

func (account *Account) ProjectsListUrl() string {
	return account.ProjectsUrl
}
