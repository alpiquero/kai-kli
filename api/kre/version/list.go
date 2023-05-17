package version

// Version represents a Version entity in KRE.
type Version struct {
	Name   string
	Status string
}

// List contains a list of  Version.
type List []Version

// List calls to KRE API and returns a list of Version entities.
func (c *versionClient) List(product string) (List, error) {
	query := `
		query GetVersions($productId: ID!) {
			versions(productId: $productId) {
				name
				status
			}
		}
	`

	vars := map[string]interface{}{
		"productId": product,
	}

	var respData struct {
		Versions List
	}

	err := c.client.MakeRequest(query, vars, &respData)

	return respData.Versions, err
}
