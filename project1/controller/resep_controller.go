package controllers

import (
	"fmt"
	"io/ioutil"
	constants "myproject/contens"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CallApi(c echo.Context) error {

	url := constants.BaseUrl + "/recipes/findByIngredients"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", "43ab617e49b24debb876ea1043a2b3f7")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return c.JSON(http.StatusOK, "")
}
