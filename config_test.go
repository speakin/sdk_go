/*
 * Copyright (c) 2018. The speakin Authors. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */
package voiceprint_cloud_go_git

type Config struct {
	App struct {
		Name      string `yaml:"name"`
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
	} `yaml:"app"`
	Bucket struct {
		Name      string `yaml:"name"`
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
	} `yaml:"bucket"`
	TestData struct {
		Users []struct {
			Name  string   `yaml:"Name"`
			Files []string `yaml:"Files"`
		} `yaml:"Users"`
		SimpleRate string `yaml:"simpleRate"`
	} `yaml:"TestData"`
}
