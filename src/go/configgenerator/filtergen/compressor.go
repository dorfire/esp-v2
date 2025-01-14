// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filtergen

import (
	"fmt"

	ci "github.com/GoogleCloudPlatform/esp-v2/src/go/configinfo"
	"github.com/GoogleCloudPlatform/esp-v2/src/go/util/httppattern"
	corepb "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	brpb "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/brotli/compressor/v3"
	gzippb "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/gzip/compressor/v3"
	comppb "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/compressor/v3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type CompressorType int

const (
	GzipCompressor CompressorType = iota
	BrotliCompressor
)

const (
	// EnvoyCompressorFilterName is the Envoy filter name for debug logging.
	EnvoyCompressorFilterName = "envoy.filters.http.compressor"

	// EnvoyBrotliCompressorName is a compressor extension name.
	EnvoyBrotliCompressorName = "envoy.compression.brotli.compressor"

	// EnvoyGzipCompressorName is a compressor extension name.
	EnvoyGzipCompressorName = "envoy.compression.gzip.compressor"
)

type CompressorGenerator struct {
	compressorType CompressorType

	// skipFilter indicates if this filter is disabled based on options and config.
	skipFilter bool
}

// NewCompressorGenerator creates the CompressorGenerator with cached config.
func NewCompressorGenerator(serviceInfo *ci.ServiceInfo, compressorType CompressorType) *CompressorGenerator {
	return &CompressorGenerator{
		compressorType: compressorType,
		skipFilter:     !serviceInfo.Options.EnableResponseCompression,
	}
}

func (g *CompressorGenerator) FilterName() string {
	return EnvoyCompressorFilterName
}

func (g *CompressorGenerator) IsEnabled() bool {
	return !g.skipFilter
}

func (g *CompressorGenerator) GenFilterConfig(serviceInfo *ci.ServiceInfo) (proto.Message, error) {
	cfg, name, err := g.getCompressorConfig()
	if err != nil {
		return nil, err
	}
	ca, err := anypb.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("error marshaling %s Compressor config to Any: %v", name, err)
	}
	return &comppb.Compressor{
		CompressorLibrary: &corepb.TypedExtensionConfig{
			Name:        name,
			TypedConfig: ca,
		},
	}, nil
}

func (g *CompressorGenerator) GenPerRouteConfig(method *ci.MethodInfo, httpRule *httppattern.Pattern) (proto.Message, error) {
	return nil, nil
}

func (g *CompressorGenerator) getCompressorConfig() (proto.Message, string, error) {
	switch g.compressorType {
	case GzipCompressor:
		return &gzippb.Gzip{}, EnvoyGzipCompressorName, nil
	case BrotliCompressor:
		return &brpb.Brotli{}, EnvoyBrotliCompressorName, nil
	}
	return nil, "", fmt.Errorf("unknown compressor type: %v", g.compressorType)
}
