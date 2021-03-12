/*
 * Copyright 2021 Luke Whrit

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *     http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

var (
	sites []string
	c     = colly.NewCollector(
		colly.MaxDepth(4),
		colly.Async(),
	)
	allowedProtocols = []string{
		"http",
		"https",
	}
)

func getSearchSites() ([]string, error) {
	var result []string

	filePath := "configs/sites.json"
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Crawl uses Colly and crawls
func Crawl() {
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 8})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		for _, protocol := range allowedProtocols {
			if strings.HasPrefix(link, protocol) {
				sites = append(sites, link)
			} else {
				return
			}
		}

		e.Request.Visit(link)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL)
	})

	searchSites, err := getSearchSites()

	if err != nil {
		log.Fatal(err)
	}

	for _, site := range searchSites {
		c.Visit(site)
	}

	// Just writing to JSON file for now for testing purposes
	defer func() {
		blob, err := json.Marshal(sites)

		if err != nil {
			log.Fatal(err)
		}

		if err = ioutil.WriteFile("./blob.json", blob, 0644); err != nil {
			log.Fatal(err)
		}
	}()

	c.Wait()
}
