package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func assetError(err error) {
	if err != nil {
		log.Fatalf("ERROR : %s", err)
	}
}

func download(urls, dir, contentStr string) string {
	var err error
	var rep *http.Response
	fmt.Println(urls)
	err = os.MkdirAll(dir, os.ModePerm)
	assetError(err)
	URL, _ := url.ParseRequestURI(urls)
	_, file := filepath.Split(URL.Path)
	rep, err = http.Get(urls)
	bytes, _ := ioutil.ReadAll(rep.Body)
	err = ioutil.WriteFile(dir+file, bytes, os.ModePerm)
	assetError(err)
	return strings.Replace(string(contentStr), urls, dir+file, 1)
}

func main() {
	rep, error := http.Get("https://studygolang.com/resources/13880")
	assetError(error)
	content, err := ioutil.ReadAll(rep.Body)
	assetError(err)

	pngReg := regexp.MustCompile(`((ht|f)tps?)://[w]{0,3}.\S+\.png`)

	jsReg := regexp.MustCompile(`((ht|f)tps?)://[w]{0,3}.\S+\.js`)

	jpgReg := regexp.MustCompile(`((ht|f)tps?)://[w]{0,3}.\S+\.jpg`)

	gifReg := regexp.MustCompile(`((ht|f)tps?)://[w]{0,3}.\S+\.gif`)

	cssReg := regexp.MustCompile(`((ht|f)tps?)://[w]{0,3}.\S+\.css`)

	contentStr := string(content)

	fmt.Println("PARSE")
	imagesDir := "/home/huizhou/桌面/goProjects/src/pac/image/png/"
	for _, d := range pngReg.FindAllString(string(contentStr), -1) {
		contentStr = download(d, imagesDir, contentStr)
	}
	jsDir := "/home/huizhou/桌面/goProjects/src/pac/js/"
	for _, d := range jsReg.FindAllString(string(content), -1) {
		contentStr = download(d, jsDir, contentStr)
	}
	jpgDir := "/home/huizhou/桌面/goProjects/src/pac/image/jpg/"
	for _, d := range jpgReg.FindAllString(string(content), -1) {
		contentStr = download(d, jpgDir, contentStr)
	}
	gifDir := "/home/huizhou/桌面/goProjects/src/pac/image/gif/"
	for _, d := range gifReg.FindAllString(string(content), -1) {
		contentStr = download(d, gifDir, contentStr)
	}

	cssDir := "/home/huizhou/桌面/goProjects/src/pac/style/css/"
	for _, d := range cssReg.FindAllString(string(content), -1) {
		contentStr = download(d, cssDir, contentStr)
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(contentStr))
	doc.Find("a").Each(func(i int, ele *goquery.Selection) {
		herf, exists := ele.Attr("herf")
		if exists {
			fmt.Println(herf)
		}

	})
	doc.Find("img").Each(func(i int, ele *goquery.Selection) {
		src, exsits := ele.Attr("src")
		if exsits {
			fmt.Println(src)
			ele.SetAttr("src", "https://studygolang.com/"+src)
		}
	})
	contentStr, _ = doc.Html()
	ioutil.WriteFile("index.html", []byte(contentStr), os.ModePerm)

}
