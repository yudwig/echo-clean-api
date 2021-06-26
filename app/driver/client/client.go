package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetAllUsers() {
	res, _ := http.Get("http://127.0.0.1:8000/user/")
	outputJson(res)
}

func CreateUser(name string) {
	values := url.Values{}
	values.Add("name", name)
	res, _ := http.PostForm("http://127.0.0.1:8000/user/", values)
	outputJson(res)
}

func UpdateUser(id string, name string) {
	type Body struct {
		Name string `json:"name"`
	}
	body, _ := json.Marshal(&Body{Name: name})
	req, _ := http.NewRequest("PATCH", "http://127.0.0.1:8000/user/"+id, bytes.NewReader(body))
	c := new(http.Client)
	res, _ := c.Do(req)
	outputJson(res)
}

func DeleteUser(id string) {
	req, _ := http.NewRequest("DELETE", "http://127.0.0.1:8000/user/"+id, nil)
	c := new(http.Client)
	res, _ := c.Do(req)
	outputJson(res)
}

func outputJson(res *http.Response) {
	body, _ := ioutil.ReadAll(res.Body)
	var buf bytes.Buffer
	_ = json.Indent(&buf, body, "", "  ")
	fmt.Println(buf.String())
}
