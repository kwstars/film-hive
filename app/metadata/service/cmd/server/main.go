package main

import (
	"flag"
	"fmt"
	"os"

	nacosKratos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/kwstars/film-hive/app/metadata/service/internal/conf"
	nacosClients "github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	_ "go.uber.org/automaxprocs"
	"gopkg.in/yaml.v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "metadata"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf    string
	flagVersion bool
	CompileTime string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
	flag.BoolVar(&flagVersion, "version", false, "if true, print version and exit")
}

func newApp(logger log.Logger, rr registry.Registrar, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(rr),
	)
}

// setTracerProvider 设置trace
func setTracerProvider(c *conf.Trace) (tp *tracesdk.TracerProvider, err error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Endpoint)))
	if err != nil {
		return nil, err
	}

	tp = tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("env", "dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	return
}

// NewNacosConfigSource 创建一个远程配置源 - Nacos
func newNacosConfigSource(configAddr string, configPort uint64, group string) config.Source {
	sc := []nacosConstant.ServerConfig{
		*nacosConstant.NewServerConfig(configAddr, configPort),
	}

	cc := nacosConstant.ClientConfig{
		// NamespaceId: "",    // 命名空间
		// Username: "",
		// Password: "",
		TimeoutMs:            10 * 1000, // http请求超时时间，单位毫秒
		BeatInterval:         5 * 1000,  // 心跳间隔时间，单位毫秒
		UpdateThreadNum:      20,        // 更新服务的线程数
		LogLevel:             "debug",
		CacheDir:             "/tmp/nacos/cache", // 缓存目录
		LogDir:               "/tmp/nacos/log",   // 日志目录
		NotLoadCacheAtStart:  true,               // 在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true,               // 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	nacosClient, err := nacosClients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	return nacosKratos.NewConfigSource(nacosClient,
		nacosKratos.WithGroup(group),
		nacosKratos.WithDataID("bootstrap.yaml"),
	)
}

func main() {
	flag.Parse()
	var (
		bc  conf.Bootstrap
		tp  *tracesdk.TracerProvider
		err error
	)

	if flagVersion {
		fmt.Println(Version, CompileTime)
		os.Exit(0)
	}

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
			newNacosConfigSource("127.0.0.1", 8848, "test"),
			// env.NewSource()
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err = c.Load(); err != nil {
		panic(err)
	}

	if err = c.Scan(&bc); err != nil {
		panic(err)
	}

	tp, err = setTracerProvider(bc.Trace)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(&bc, tp)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
