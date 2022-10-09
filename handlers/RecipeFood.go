package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "foodRe/helpers"
	. "foodRe/models"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

func RecipeFood(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("image")
	authForLogMeal := r.FormValue("apiKeyLogMeal")

	if err != nil {
		CheckError(err)
		return
	}

	defer file.Close()

	tempFile, err := ioutil.TempFile("../uploads", "upload-*.jpg")
	if err != nil {
		CheckError(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		CheckError(err)
		return
	}
	tempFile.Write(fileBytes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var foodData RecognizedFood

	result := calls("https://api.logmeal.es/v2/image/segmentation/complete/v1.0", "POST", tempFile.Name(), authForLogMeal)

	if err := json.Unmarshal(result, &foodData); err != nil {
		panic(err)
	}
	ttt := foodData.SegmentationResults[0]

	Count := make(map[int]string)
	Count2 := make(map[int]float64)
	for k, v1 := range ttt.RecognitionResults {

		Count[k] = v1.Name
		Count2[k] = v1.Prob

	}
	foodName := Count[0]

	w.Write(callTheMealDb(foodName))

}

func call(urlPath, method, namesFile, authForLogMeal string) []byte {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("image", namesFile)
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Open(namesFile)
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		fmt.Println(err)
	}
	writer.Close()
	req, err := http.NewRequest(method, urlPath, bytes.NewReader(body.Bytes()))
	if err != nil {
		fmt.Println(err)
	}
	//d6f3fc37780263a9b79b3882f98c45e926aefe6f
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+authForLogMeal+"")
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	//body, err := ioutil.ReadAll(res.Body)

	defer rsp.Body.Close()
	bodyss, _ := ioutil.ReadAll(rsp.Body) // response body is []byte
	return bodyss

}

func callTheMealDb(foodName string) []byte {

	foodName = strings.ReplaceAll(foodName, " ", "%20")

	url := "https://www.themealdb.com/api/json/v1/1/search.php?s=" + foodName + ""
	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}
	return body

}
