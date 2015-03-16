package api

import (
	"fmt"
)

type Project struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	ApiKey        string   `json:"api_key"`
	Errors        int      `json:"errors"`
	Icon          string   `json:"icon"`
	ReleaseStages []string `json:"release_stages"`
	Type          string   `json:"type"`
	CreatedAt     jTime    `json:"created_at"`
	UpdatedAt     jTime    `json:"updated_at"`
	ErrorsUrl     string   `json:"errors_url"`
	EventsUrl     string   `json:"events_url"`
	HtmlUrl       string   `json:"html_url"`
	Url           string   `json:"url"`
}

func (c *Connection) Projects(object ProjectsList, params Parameters) ([]Project, error) {
	var resp []Project
	paramString, err := params.paramString()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s", object.ProjectsListUrl(), paramString)
	err = c.client.Query("GET", url, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connection) ProjectById(projectId string) (Project, error) {
	url := fmt.Sprintf("/projects/%s", projectId)
	var resp Project
	err := c.client.Query("GET", url, &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Connection) NewProject(account *Account, name string, projecttype string) (Project, error) {
	url := fmt.Sprintf("/accounts/%s/projects", account.Id)
	var resp Project
	newProject := EditProject{Name: name, Type: projecttype}
	err := c.client.Exec("POST", url, newProject, &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Connection) UpdateProject(project *Project, name string, projecttype string) (Project, error) {
	url := fmt.Sprintf("/projects/%s", project.Id)
	var resp Project
	editProject := EditProject{Name: name, Type: projecttype}
	err := c.client.Exec("PATCH", url, editProject, &resp)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Connection) DeleteProject(project *Project) error {
	url := fmt.Sprintf("/projects/%s", project.Id)
	noBody := NoBody{}
	err := c.client.Exec("DELETE", url, noBody, noBody)

	if err != nil {
		return err
	}

	return nil
}

func (project *Project) UserListUrl() string {
	url := fmt.Sprintf("%s/projects/%s/users", bugsnagApi, project.Id)
	return url
}

type ProjectsList interface {
	ProjectsListUrl() string
}

type EditProject struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
