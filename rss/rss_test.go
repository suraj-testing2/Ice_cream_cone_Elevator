/*
Copyright 2013 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package rss

import (
	"io/ioutil"
	"testing"
)

func assert(t *testing.T, expected string, actual string, description string) {
	if actual != expected {
		t.Errorf("%v expected: '%v', actual: '%v'", description, expected, actual)
	}
}

func TestParse(t *testing.T) {
	xml, err := ioutil.ReadFile("testdata/rss2sample.xml")
	if err != nil {
		t.Errorf("%v", err)
	}
	doc, err := Parse(xml)
	if err != nil {
		t.Errorf("%v", err)
	}
	assert(t, "Liftoff News", doc.Title, "Wrong title")
	assert(t, "http://liftoff.msfc.nasa.gov/", doc.Link, "Wrong link")
}
