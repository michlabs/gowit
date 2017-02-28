package gowit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"net/url"
	"strings"
)

const (
	APIEndpoint = "https://api.wit.ai"
	APIVersion = "20170121"
	
	TraitLookup = "trait"
	KeywordsLookup = "keywords"
)

var apiToken string

// Client represents a client for the Wit.AI API
type Client struct {}

type param struct {
	Method string
	Path string
	ContentType string
	Data []byte
}

func NewClient(token string) *Client {
	apiToken = token
	return &Client{}
}

func (c *Client) Detect(text string) (Meaning, error) {
	var p param
	p.Method = "GET"
	p.Path = "/message?q=" + url.QueryEscape(text)

	var m Meaning
	res, err := request(&p)
	if err != nil {
		return m, err
	}

	if err := json.Unmarshal(res, &m); err != nil {
		fmt.Println("Failed to unmarshal response: ", string(res))
		return m, err
	}

	return m, nil
}

// ListEntities returns a list of available entities for the app
// TODO: Should we return []Entity or just []string?
func (c *Client) ListEntities() ([]Entity, error) {
	var p param
	p.Method = "GET"
	p.Path = "/entities"
	
	res, err := request(&p)
	if err != nil {
		return nil, err
	}

	var entityNames []string
	if err := json.Unmarshal(res, &entityNames); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal: %s. Data: %s", err.Error(), string(res))
	}

	var entities []Entity
	for _, n := range entityNames {
		var e Entity
		e.Name = n
		entities = append(entities, e)
	}

	return entities, nil
}

// GetEntity returns all the expressions validated for an entity. 
// Wit.AI currently limits to the first 1000 values (with the first 50 expressions)
func (c *Client) GetEntity(id string) (e Entity, err error) {
	var p param
	p.Method = "GET"
	p.Path = "/entities/" + id 

	res, err := request(&p)
	if err != nil {
		return e, err
	}
	return parseEntity(res)
}

// UpdateEntity updates an entity with the given attributes.
func (c *Client) UpdateEntity(e *Entity) error {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal entity: %s", err.Error())
	}

	var p param
	p.Method = "PUT"
	p.Path = "/entities/" + e.Name
	p.Data = data

	_, err = request(&p)
	return err
}

// func (c *Client) DeleteAllValue

func deleteValue(e *Entity, v *Value) error {
	var p param 
	p.Method = "DELETE"
	p.Path = "/entities/" + e.Name + "/values/" + v.Name
	_, err := request(&p)
	return err
}

func parseEntity(data []byte) (Entity, error) {
	var e Entity 
	if err := json.Unmarshal(data, &e); err != nil {
		return e, fmt.Errorf("Failed to unmarshal: %s. Data: %s", err.Error(), string(data))
	}
	return e, nil
}

func request(p *param) ([]byte, error) {
	if strings.Contains(p.Path, `?`) {
		p.Path += "&v=" + APIVersion
	} else {
		p.Path += "?v=" + APIVersion
	}

	req, err := http.NewRequest(p.Method, APIEndpoint + p.Path, bytes.NewReader(p.Data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + apiToken)
	req.Header.Set("Accept", "application/json")
	if p.ContentType != "" {
		req.Header.Set("Content-Type", p.ContentType)
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(http.StatusText(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return body, nil
}