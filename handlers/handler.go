package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetParam(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	param, ok := vars["PARAM"]
	if !ok {
		fmt.Println("Failed")
	}
	res.Write([]byte(fmt.Sprintf("Hello, %s", param)))
}

func GetBad(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "", http.StatusInternalServerError)
}

func PostBody(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.Write([]byte(err.Error()))
	}

	returnData := append([]byte("I got a message: \n"), data...)

	res.Write(returnData)
}

func PostHeaders(res http.ResponseWriter, req *http.Request) {
	a, _ := strconv.Atoi(req.Header.Get("a"))
	b, _ := strconv.Atoi(req.Header.Get("b"))

	res.Header().Set("a+b", strconv.Itoa(a+b))
}
