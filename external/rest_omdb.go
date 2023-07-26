package external

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetMovieByID(id string) (err error, resBody []byte) {
	requestURL := "https://www.omdbapi.com/?apikey=faf7e5bb&i=" + id
	//requestURL := os.Getenv("PROVIDER_URL") + "/?apikey=" + os.Getenv("API_KEY") + "&i=" + id

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("GetMovieByID - client: could not create request: %s\n", err)
		return err, nil
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("GetMovieByID - client: error making http request: %s\n", err)
		return err, nil
	}

	fmt.Printf("GetMovieByID - client: got response!\n")
	fmt.Printf("GetMovieByID - client: status code: %d\n", res.StatusCode)

	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("GetMovieByID - client: could not read response body: %s\n", err)
		return err, nil
	}
	fmt.Printf("GetMovieByID - client: response body: %s\n", resBody)
	return nil, resBody
}

func SearchMovie(query string, movieType string, page int) (err error, resBody []byte) {
	requestURL := "http://www.omdbapi.com/?apikey=faf7e5bb&s=" + query + "&page=" + strconv.Itoa(page) + "&type=" + movieType
	//requestURL := "http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("SearchMovie - client: could not create request: %s\n", err)
		return err, nil
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("SearchMovie - client: error making http request: %s\n", err)
		return err, nil
	}

	fmt.Printf("SearchMovie - client: got response!\n")
	fmt.Printf("SearchMovie - client: status code: %d\n", res.StatusCode)

	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("SearchMovie - client: could not read response body: %s\n", err)
		return err, nil
	}
	fmt.Printf("SearchMovie - client: response body: %s\n", resBody)
	return nil, resBody
}
