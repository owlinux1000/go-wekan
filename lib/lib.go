package lib

import (
    "fmt"
    "net/url"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

const ENDPOINT = "https://wekan.alicemacs.com"

type LoginResponse struct {
    Id string `json:"id,omitempty"`
    Token string `json:"token,omitempty"`
    TokenExpires string `json:"tokenExpires,omitempty"`
    Error interface{} `json:"error,omitempty"`
    Reason string `json:"reason,omitempty"`
}

func Register(username string, password string, email string) LoginResponse {
    
    values := url.Values{}
    values.Set("username", username)
    values.Add("password", password)
    values.Add("email", email)
    
    resp, err := http.PostForm(ENDPOINT + "/users/register", values)
    if err != nil {
        fmt.Println(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    if err != nil {
        fmt.Println(err.Error())
    }

    login_resp := LoginResponse{}
    err = json.Unmarshal(body, &login_resp)
    if err != nil {
        fmt.Println(err.Error())
        return login_resp
    }

    return login_resp
}

func Login(username string, password string) LoginResponse {

    values := url.Values{}
    values.Set("username", username)
    values.Add("password", password)
    
    resp, err := http.PostForm(ENDPOINT + "/users/login", values)
    if err != nil {
        fmt.Println(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    if err != nil {
        fmt.Println(err.Error())
    }

    login_resp := LoginResponse{}
    err = json.Unmarshal(body, &login_resp)
    if err != nil {
        fmt.Println(err.Error())
        return login_resp
    }

    return login_resp
    
}
