package dodo

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Settings struct {
	URL string
	Username string
	Password string
	Salt string
	JWTSecret string
}

type Dodo struct {
	settings Settings
	usingAuth bool
}

func NewDodoConnection(st Settings) (*Dodo, error) {
	d := new(Dodo)

	_, err := http.Get(st.URL)
	if err != nil {
		return nil, errors.New("Could not connect to Dodo with given URL.")
	}

	// If any of these are empty, assuming no auth
	d.usingAuth = true
	if st.Username == "" || st.Password == "" || st.Salt == "" || st.JWTSecret == "" {
		d.usingAuth = false
	}

	d.settings = st

	return d, nil
}

func (d *Dodo) Store(doc map[string]interface{}) (string, error) {
	j, err := json.Marshal(doc)
    if err != nil {
        return "", err
    }
	
	c := &http.Client{}
	r, err := http.NewRequest(http.MethodPut, d.settings.URL, bytes.NewReader(j))
	if err != nil {
        return "", err
	}
	rs, err := c.Do(r)
	if err != nil {
        return "", err
    }

	defer rs.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(rs.Body).Decode(&result)

	return result["id"].(string), nil
}

func (d *Dodo) Get(id string) (map[string]interface{}, error) {
	r, err := http.Get(d.settings.URL + "/document/" + id)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)

	return result, nil
}

func (d *Dodo) Modify(id string, doc map[string]interface{}) error {
	j, err := json.Marshal(doc)
    if err != nil {
        return err
	}
	
	_, err = http.Post(d.settings.URL + "/document/" + id, "application/json", bytes.NewReader(j))
	if err != nil {
		return err
	}
	
	return nil
}

func (d *Dodo) GetAll() ([]map[string]interface{}, error) {
	r, err := http.Get(d.settings.URL + "/all")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var result []map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)

	return result, nil
}

func (d *Dodo) Delete(id string) error {
	c := &http.Client{}
	r, err := http.NewRequest(http.MethodDelete, d.settings.URL + "/document/" + id, nil)
	if err != nil {
        return err
	}
	_, err = c.Do(r)
	if err != nil {
        return err
    }
	
	return nil
}

func (d *Dodo) DeleteAll(id string) error {
	c := &http.Client{}
	r, err := http.NewRequest(http.MethodDelete, d.settings.URL + "/all", nil)
	if err != nil {
        return err
	}
	_, err = c.Do(r)
	if err != nil {
        return err
    }
	
	return nil
}
