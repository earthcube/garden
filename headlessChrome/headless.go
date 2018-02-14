package main

// This code needs docker image knqz/chrome-headless which you can get with
//  docker pull knqz/chrome-headless
// run with
//  docker run -d -p 9222:9222 --rm --name chrome-headless knqz/chrome-headless

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
)

func main() {
	var err error

	// // Set up our log file for runs...
	// f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()
	// log.SetOutput(f)

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome
	c, err := chromedp.New(ctxt, chromedp.WithTargets(client.New().WatchPageTargets(ctxt)), chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var site, res string
	// err = c.Run(ctxt, text("https://datanext-mmldquydwe.now.sh/", &res)) // Example page Bryce set up at DataOne doing dynamic DOM update
	// err = c.Run(ctxt, text("https://handy-owl.nceas.ucsb.edu/view/test.10.1", &res)) // Example page Bryce set up at DataOne doing dynamic DOM update
	err = c.Run(ctxt, text("https://www2.earthref.org/MagIC/16403", &res))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Text %s --> %s", site, res)
}

// You should be able to use XPath to select that node and retrieve the contents of the node (element).
// /html/body/section/article
func text(targeturl string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(targeturl),
		// chromedp.Text(`#loading-app`, res, chromedp.NodeVisible, chromedp.ByID),
		// chromedp.Text(`script[type=application/ld+json]`, res, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`/html/head/title`, res, chromedp.NodeVisible, chromedp.ByQuery),
		// chromedp.Text(`span`, res, chromedp.NodeVisible, chromedp.ByQuery),
	}
}
