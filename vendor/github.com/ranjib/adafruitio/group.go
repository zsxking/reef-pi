package adafruitio

type Group struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Feeds       []Feed `json:"feeds"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func (c *Client) ListGroups(u string) ([]Group, error) {
	var groups []Group
	resp, err := c.get("/" + u + "/groups")
	if err != nil {
		return groups, err
	}
	return groups, c.decodeJSON(resp, &groups)
}

func (c *Client) CreateGroup(u string, g Group) error {
	_, err := c.post("/"+u+"/groups", g)
	return err
}

func (c *Client) DeleteGroup(u, g string) error {
	_, err := c.delete("/" + u + "/groups/" + g)
	return err
}

func (c *Client) GetGroup(u, g string) (Group, error) {
	var group Group
	resp, err := c.get("/" + u + "/groups/" + g)
	if err != nil {
		return group, err
	}
	return group, c.decodeJSON(resp, &group)
}

func (c *Client) UpdateGroup(u, g string, group Group) error {
	_, err := c.patch("/"+u+"/groups/"+g, group)
	return err
}

func (c *Client) ReplaceGroup(u, g string, group Group) error {
	_, err := c.put("/"+u+"/groups/"+g, group)
	return err
}
