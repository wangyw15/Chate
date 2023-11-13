package service

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/wangyw15/Chate/util"
)

func GetImageUrl(jan string) string {
	resp, err := util.HttpClient.Get("https://www.animate-onlineshop.jp/products/list.php?smt=" + jan + "&nf=1")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ""
	}

	imgurl, success := doc.Find("div.item_list li img").Attr("src")
	if !success {
		return ""
	}
	parsed, err := url.Parse(imgurl)
	if err != nil {
		return ""
	}
	return parsed.Scheme + "://" + parsed.Host + parsed.Path + "?image=" + parsed.Query()["image"][0]
}
