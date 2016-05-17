package main

/**
* a package to download movie url of www.qtfy30.cn
 */

import (
	//"movie/downloader"
	"fmt"
)

var urlTemplete = "http://www.qtfy30.cn/page/%d/"

func main() {
	InitHtmlTag("videoInfo.html")
	for i := 1; i < 234; i++ {
		listUrl := fmt.Sprintf(urlTemplete, i)
		FindAllMovieInfoPage(listUrl)
	}
}
