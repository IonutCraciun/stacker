package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/anuvu/stacker"
	"github.com/anuvu/stacker/container"
	"github.com/anuvu/stacker/types"
)

var buildCmd = cli.Command{
	Name:   "build",
	Usage:  "builds a new OCI image from a stacker yaml file",
	Action: doBuild,
	Flags:  initBuildFlags(),
	Before: beforeBuild,
}

func initBuildFlags() []cli.Flag {
	return append(
		initCommonBuildFlags(),
		cli.StringFlag{
			Name:  "stacker-file, f",
			Usage: "the input stackerfile",
			Value: "stacker.yaml",
		})
}

func initCommonBuildFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "leave-unladen",
			Usage: "leave the built rootfs mount after image building",
		},
		cli.BoolFlag{
			Name:  "no-cache",
			Usage: "don't use the previous build cache",
		},
		cli.StringSliceFlag{
			Name:  "substitute",
			Usage: "variable substitution in stackerfiles, FOO=bar format",
		},
		cli.StringFlag{
			Name:  "on-run-failure",
			Usage: "command to run inside container if run fails (useful for inspection)",
		},
		cli.BoolFlag{
			Name:  "shell-fail",
			Usage: "exec /bin/sh inside the container if run fails (alias for --on-run-failure=/bin/sh)",
		},
		cli.StringSliceFlag{
			Name:  "layer-type",
			Usage: "set the output layer type (supported values: tar, squashfs); can be supplied multiple times",
			Value: &cli.StringSlice{"tar"},
		},
		cli.BoolFlag{
			Name:  "order-only",
			Usage: "show the build order without running the actual build",
		},
	}
}

func beforeBuild(ctx *cli.Context) error {
	if config.StorageType == "overlay" && ctx.Bool("leave-unladen") {
		return errors.Errorf("cannot use --storage-type=overlay and --leave-unladen together")
	}

	// Validate build failure arguments
	err := validateBuildFailureFlags(ctx)
	if err != nil {
		return err
	}

	// Validate layer type
	err = validateLayerTypeFlags(ctx)
	if err != nil {
		return err
	}
	return nil
}

func newBuildArgs(ctx *cli.Context) (stacker.BuildArgs, error) {
	args := stacker.BuildArgs{
		Config:       config,
		LeaveUnladen: ctx.Bool("leave-unladen"),
		NoCache:      ctx.Bool("no-cache"),
		Substitute:   ctx.StringSlice("substitute"),
		OnRunFailure: ctx.String("on-run-failure"),
		OrderOnly:    ctx.Bool("order-only"),
		Progress:     shouldShowProgress(ctx),
	}
	var err error
	args.LayerTypes, err = types.NewLayerTypes(ctx.StringSlice("layer-type"))
	return args, err
}

func doBuild(ctx *cli.Context) error {
	args, err := newBuildArgs(ctx)
	if err != nil {
		return err
	}

	// if it's false run stacker in userns
	fmt.Println("build --internal-userns:", args.BuildinUserNS, " ctx.Bool(internal-userns)", ctx.Bool("internal-userns")) // default is false
	fmt.Println("args.config.userns ", args.Config.Userns)
	if !args.Config.Userns {
		binary, err := os.Readlink("/proc/self/exe")
		if err != nil {
			return err
		}
		err = container.MaybeRunInUserns([]string{binary, "--internal-userns", "build"}, "message")
		//err := container.MaybeRunInUserns([]string{"ls", "-al"}, "string")
		//err := container.RunInUserns2([]string{"--internal-userns", "build"}, "message")
		return err
	}

	builder := stacker.NewBuilder(&args)
	fmt.Println("argss")
	fmt.Println(args)
	// fmt.Println("builder")
	// fmt.Println(builder)
	fmt.Println(ctx.String("stacker-file"))
	return builder.BuildMultiple([]string{ctx.String("stacker-file")})
}
