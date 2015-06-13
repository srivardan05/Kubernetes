/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package aws_cloud

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	"github.com/aws/aws-sdk-go/aws"
)

func stringSetToPointers(in util.StringSet) []*string {
	if in == nil {
		return nil
	}
	out := make([]*string, len(in))
	for k := range in {
		out = append(out, aws.String(k))
	}
	return out
}

func stringSetFromPointers(in []*string) util.StringSet {
	if in == nil {
		return nil
	}
	out := util.NewStringSet()
	for i := range in {
		out.Insert(orEmpty(in[i]))
	}
	return out
}

func orZero(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}
