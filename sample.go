// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"strings"
)

// sample represents a Google Cloud Platform sample and associated properties.
type sample struct {
	// The Google Cloud Project ID this sample will deploy to
	googleCloudProject string

	// The local directory this sample is located in
	dir string

	// The cloudRunService this sample will deploy to
	service *cloudRunService

	// The cloudContainerImage that represents the location of this sample's build container image in the GCP Container
	// Registry
	container *cloudContainerImage

	// The lifecycle for building and deploying this sample to Cloud Run
	buildDeployLifecycle *lifecycle
}

// newSample creates a new sample object for the sample located in the provided local directory.
func newSample(dir string) *sample {
	s := sample{
		googleCloudProject: os.Getenv("GOOGLE_CLOUD_PROJECT"),
		dir:                dir,
	}

	s.service = newCloudRunService(&s)
	s.container = newCloudContainerImage(&s)
	s.buildDeployLifecycle = getLifecycle(&s)

	return &s
}

// sampleName computes a sample name for a sample object. Right now, it's defined as a shortened version of the sample's
// local directory. Its length is flexible based on the provided length of a suffix that will be appended to the end of
// the name.
func (s *sample) sampleName(suffixLen int) string {
	n := strings.ReplaceAll(s.dir[len(s.dir)-(maxCloudRunServiceNameLen-suffixLen):], "/", "-")

	if n[len(n)-1] == '-' {
		n = n[:len(n)-1]
	}

	if n[0] == '-' {
		n = n[1:]
	}

	return strings.ToLower(n)
}
