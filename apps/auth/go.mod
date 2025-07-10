module github.com/yclw/mys_project/apps/auth

go 1.24.2

require (
	github.com/spf13/viper v1.20.1
	github.com/yclw/mys_project/pkg/common v0.0.0-00010101000000-000000000000
	github.com/yclw/mys_project/pkg/protobuf v0.0.0-00010101000000-000000000000
	github.com/yclw/mys_project/pkg/utils v0.0.0-20250710043614-eb9747846d0c
	google.golang.org/grpc v1.73.0
)

replace (
	github.com/yclw/mys_project/pkg/common => ../../pkg/common
	github.com/yclw/mys_project/pkg/model => ../../pkg/model
	github.com/yclw/mys_project/pkg/protobuf => ../../pkg/protobuf
	github.com/yclw/mys_project/pkg/utils => ../../pkg/utils
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/sagikazarmark/locafero v0.7.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.12.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.etcd.io/etcd/api/v3 v3.6.2 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.6.2 // indirect
	go.etcd.io/etcd/client/v3 v3.6.2 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
