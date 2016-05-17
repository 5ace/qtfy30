package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/opesun/goquery"
	. "movie/util"
	"strings"
	"time"
)

var RETRY_TIME = 4

var movieHtmlTemplete = "<a href=\"%s\">%s</a><br>\n"
var indextmlTemplete = "<a href=\"./pages/%s.html\">%s</a><br>\n"

func Test() {
	url := "http://www.qtfy30.cn/page/229/"
	p, err := goquery.ParseUrl(url)
	if err != nil {
		panic(err)
	} else {
		//pTitle := p.Find("titleimg").Html() //直接提取title的内容
		fmt.Println(p.Find(".titleimg a").Attrs("href"))

	}
}

func FindAllMovieInfoPage(url string) {
	p, err := goquery.ParseUrl(url)
	if err != nil {
		retry := 0
		for err != nil && retry < RETRY_TIME {
			time.Sleep(20 * time.Second)
			p, err = goquery.ParseUrl(url)
			if err != nil {
				retry++
				postion := "download " + url + ": \n"
				errStr := err.Error() + "\n"
				WriteInfoFile(postion, "log")
				WriteInfoFile(errStr, "log")
			}
		}
	} else {
		//pTitle := p.Find("titleimg").Html() //直接提取title的内容
		//movieTitle := strings.TrimSpace(p.Find(".mecctitle").Text())
		//fmt.Println(movieTitle)
		pageList := p.Find(".titleimg a")

		for i := 0; i < pageList.Length(); i++ {
			pageUrl := pageList.Eq(i).Attr("href")
			go GetVideoDownloadUrlFromMoviePage(pageUrl)
			time.Sleep(3 * time.Second)
		}
	}
	fmt.Println("finish process page list -- : " + url + "\n")

}

func GetVideoDownloadUrlFromMoviePage(url string) {
	p, err := goquery.ParseUrl(url)
	if err != nil {
		retry := 0
		for err != nil && retry < RETRY_TIME {
			time.Sleep(20 * time.Second)
			p, err = goquery.ParseUrl(url)
			if err != nil {
				retry++
				postion := "download " + url + ": \n"
				errStr := err.Error() + "\n"
				WriteInfoFile(postion, "log")
				WriteInfoFile(errStr, "log")
			}
		}
	} else {
		//pTitle := p.Find("titleimg").Html() //直接提取title的内容
		//movieTitle := strings.TrimSpace(p.Find(".mecctitle").Text())
		//fmt.Println(movieTitle)

		md5Computer := md5.New()
		md5Computer.Write([]byte(url))
		cipherStr := md5Computer.Sum(nil)
		md5PageName := hex.EncodeToString(cipherStr)

		pageTitle := p.Find("title").Text()
		pageInfo := fmt.Sprintf(movieHtmlTemplete, "pages/"+md5PageName+".html", pageTitle)
		WriteInfoFile(pageInfo, "videoInfo.html")

		pageList := p.Find(".content-text p a")

		htmlPath := "pages/" + md5PageName + ".html"

		InitHtmlTag(htmlPath)
		for i := 0; i < pageList.Length(); i++ {
			videoUrl := pageList.Eq(i).Attr("href")
			videoTitle := strings.TrimSpace(pageList.Eq(i).Text())
			videoInfo := fmt.Sprintf(movieHtmlTemplete, videoUrl, videoTitle)
			WriteInfoFile(videoInfo, htmlPath)
			time.Sleep(time.Second)
		}
		CloseHtmlTag(htmlPath)
	}
	fmt.Println("finish process page : " + url + "\n")
}

func InitHtmlTag(fileName string) {
	initHtmlStr := "<meta charset=\"UTF-8\">"
	WriteInfoFile(initHtmlStr, fileName)
}

func CloseHtmlTag(fileName string) {
	closeHtmlTag := "</html>"
	WriteInfoFile(closeHtmlTag, fileName)

}
