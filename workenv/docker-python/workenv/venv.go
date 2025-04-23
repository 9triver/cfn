package workenv

import (
	"context"
	"github.com/9triver/cfn/work-platform/docker-python/utils"
	"github.com/9triver/cfn/work-platform/docker-python/utils/errors"
	"os"
	"os/exec"
	"sync"
	"time"
)

type VirtualEnv struct {
	mu  sync.Mutex
	ctx context.Context

	Name     string   `json:"name"`
	Exec     string   `json:"exec"`
	Packages []string `json:"packages"`
}

func (v *VirtualEnv) Interpreter() string {
	return v.Exec
}

func (v *VirtualEnv) RunPip(args ...string) (*exec.Cmd, context.CancelFunc) {
	args = append([]string{"-m", "pip"}, args...)
	cmdCtx, cancel := context.WithTimeout(v.ctx, 300*time.Second)
	cmd := exec.CommandContext(cmdCtx, v.Exec, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd, cancel
}

func (v *VirtualEnv) AddPackages(p ...string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	pkgSet := utils.MakeSetFromSlice(v.Packages)
	for _, pkg := range p {
		if pkgSet.Contains(pkg) {
			continue
		}

		if err := func() error {
			cmd, cancel := v.RunPip("install", pkg)
			defer cancel()

			if err := cmd.Run(); err != nil {
				// TODO: 把命令输出的内容返回给前端
				return errors.WrapWith(err, "venv %s: failed installing package %s", v.Name, p)
			}
			return nil
		}(); err != nil {
			return err
		}

		v.Packages = append(v.Packages, pkg)
		pkgSet.Add(pkg)
	}
	return nil
}
