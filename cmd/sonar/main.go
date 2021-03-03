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

package main

import (
	"log"

	"github.com/lukewhrit/sonar/internal/app/crawler"
	"github.com/lukewhrit/sonar/internal/app/server"
	"github.com/lukewhrit/sonar/internal/pkg/config"
)

// Code executed here should be for initialize and loading things used globally
// throughout the app like configuration
func init() {
	handleError(config.Load("configs/sonar.json"))
}

// Functions called here should power the general functionality of the app
func main() {
	crawler.Crawl()
	handleError(server.Start(
		config.Config.Server.Port,
		config.Config.Server.Hostname,
	))
}

// Simple utility function for handling errors, so that logic isn't repeated too much
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
