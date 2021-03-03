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

package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

// IConfig is the actual type for the Config structure
type IConfig struct {
	Crawler ICrawler `koanf:"crawler"`
	Server  IServer  `koanf:"server"`
}

// ICrawler represents the second-level Crawler field of IConfig
type ICrawler struct {
	SiteDepth   int `koanf:"site_depth"`
	RunInterval int `koanf:"run_interval"`
}

// IServer represents the second-level Server field of IConfig
type IServer struct {
	Port             int    `koanf:"port"`
	Hostname         string `koanf:"hostname"`
	DatabaseLocation string `koanf:"database_location"`
	MaxPayloadSize   int    `koanf:"max_payload_size"`
	Ratelimits       IRatelimits
}

// IRatelimits represents the third-level Ratelimits field of IServer
type IRatelimits struct {
	AllowedRequests int `koanf:"allowed_requests"`
	ResetDuration   int `koanf:"reset_duration"`
}

// Config is the object struct for Sonar
var (
	Config IConfig
	k      = koanf.New(".")
)

// Load loads configuration from files in the `configs/` folder to Golang object structures
func Load(path string) error {
	// Load default config values
	k.Load(structs.Provider(IConfig{
		Crawler: ICrawler{
			SiteDepth:   3,  // Dictates how deep the crawler should go when searching for links
			RunInterval: 90, // Dictates how often the crawler should run, in minutes
		},
		Server: IServer{
			Port:             3321,
			Hostname:         "0.0.0.0",
			DatabaseLocation: "/var/badger", // Path corresponding to Badger database file location
			MaxPayloadSize:   256,           // Maximum allowed size of request payloads in kilobytes
			Ratelimits: IRatelimits{
				AllowedRequests: 200,     // Allows 200 requests per 5 minutes (300k ms)
				ResetDuration:   300_000, // This value is in milliseconds
			},
		},
	}, "koanf"), nil)

	// Load JSON configuration from file at `path`
	if err := k.Load(file.Provider(path), json.Parser()); err != nil {
		return err
	}

	return k.Unmarshal("", &Config)
}
