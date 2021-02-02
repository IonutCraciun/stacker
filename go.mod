module github.com/anuvu/stacker

go 1.12

require (
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be
	github.com/apex/log v1.9.0
	github.com/cheggaaa/pb/v3 v3.0.5
	github.com/containers/image/v5 v5.9.0
	github.com/containers/libtrust v0.0.0-20200511145503-9c3a6c22cd9a // indirect
	github.com/containers/ocicrypt v1.0.3 // indirect
	github.com/containers/storage v1.24.5 // indirect
	github.com/docker/docker v20.10.2+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.10.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20200913210552-0d938eb266f3 // indirect
	github.com/freddierice/go-losetup v0.0.0-20170407175016-fc9adea44124
	github.com/google/uuid v1.1.5 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/klauspost/pgzip v1.2.5
	github.com/lxc/lxd v0.0.0-20210118115829-e1dce7da7066
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mitchellh/hashstructure v1.1.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.2-0.20190823105129-775207bd45b6
	github.com/opencontainers/umoci v0.4.7-0.20201029051143-b09d036cbfde
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0 // indirect
	github.com/prometheus/procfs v0.3.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.6.1
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/twmb/algoimpl v0.0.0-20170717182524-076353e90b94
	github.com/udhos/equalfile v0.3.0
	github.com/ulikunitz/xz v0.5.9 // indirect
	github.com/urfave/cli v1.22.5
	github.com/vbatts/go-mtree v0.5.0
	github.com/vbauerster/mpb/v5 v5.4.0 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/lxc/go-lxc.v2 v2.0.0-20210115212319-bc0fac983987
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/containers/image/v5 => github.com/anuvu/image/v5 v5.0.0-20200615203753-755940754545

replace github.com/freddierice/go-losetup => github.com/tych0/go-losetup v0.0.0-20200513233514-d9566aa43a61
