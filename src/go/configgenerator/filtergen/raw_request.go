package filtergen

import (
	"github.com/GoogleCloudPlatform/esp-v2/src/go/configinfo"
	"github.com/GoogleCloudPlatform/esp-v2/src/go/util/httppattern"
	corepb "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	luapb "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/lua/v3"
	"google.golang.org/protobuf/proto"
)

const (
	// RawRequestHeaderFilterName is the Envoy filter name for debug logging.
	RawRequestHeaderFilterName = "io.utila.espv2.filters.http.raw_req"
)

type RawRequestHeaderGenerator struct{}

func (g *RawRequestHeaderGenerator) FilterName() string {
	return RawRequestHeaderFilterName
}

func (g *RawRequestHeaderGenerator) IsEnabled() bool {
	return true
}

func (g *RawRequestHeaderGenerator) GenFilterConfig(_ *configinfo.ServiceInfo) (proto.Message, error) {
	return &luapb.Lua{
		DefaultSourceCode: &corepb.DataSource{Specifier: &corepb.DataSource_InlineString{
			InlineString: `
				function envoy_on_request(request_handle)
					local reqBuf = request_handle:body()
					if reqBuf:length() > 1024 then
						handle:logWarn("request body too large; will not copy to header")
						return
					end
					local firstKib = reqBuf:getBytes(0, 1024)
					local firstKibB64 = handle:base64Escape(firstKib)
					request_handle:headers():add("request-body-b64", firstKibB64)
				end
			`,
		}},
	}, nil
}

func (g *RawRequestHeaderGenerator) GenPerRouteConfig(_ *configinfo.MethodInfo, _ *httppattern.Pattern) (proto.Message, error) {
	return nil, nil
}
