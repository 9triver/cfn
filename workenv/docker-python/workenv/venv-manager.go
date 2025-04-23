package workenv

import (
	"context"
	"encoding/json"
	"github.com/9triver/cfn/work-platform/docker-python/utils"
	"github.com/9triver/cfn/work-platform/docker-python/utils/errors"
	"os"
	"os/exec"
	"path"
	"runtime"
	"sync"
)

const (
	venvStorageName = "actor-platform"
	venvMetadata    = ".envs.json"
	venvStart       = "__actor_executor.py"
)

var pythonExec = func() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return "python"
	default:
		return "python3"
	}
}()

var venvPath = func() string {
	home, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	dir := path.Join(home, venvStorageName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	return dir
}()

type VenvManager struct {
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
	//manager remote.ExecutorManager
	//started map[string]bool

	//SystemExec string                         `json:"system_py"`
	Envs utils.Map[string, *VirtualEnv] `json:"envs"`
}

func NewManager(ctx context.Context) (*VenvManager, error) {
	ctx, cancel := context.WithCancel(ctx)
	m := &VenvManager{
		ctx:    ctx,
		cancel: cancel,
		//manager:    manager,
		//started:    make(map[string]bool),
		//SystemExec: pythonExec,
		Envs: utils.MakeMap[string, *VirtualEnv](),
	}

	if _, err := os.Stat(venvPath); os.IsNotExist(err) {
		if err := os.MkdirAll(venvPath, os.ModePerm); err != nil {
			return nil, errors.WrapWith(err, "venv manager: error creating dir")
		}
		return m, nil
	}

	if data, err := os.ReadFile(path.Join(venvPath, venvMetadata)); err != nil {
		return m, nil
	} else if err := json.Unmarshal(data, m); err != nil {
		return nil, errors.WrapWith(err, "venv manager: error reading metadata")
	}

	for _, env := range m.Envs {
		//env.started = false
		env.ctx = ctx
		//env.handler = manager.NewExecutor(env.Name)
		//env.futures = make(map[string]utils.Future[proto.Object])
	}
	return m, nil
}

func (m *VenvManager) template() (string, bool) {
	//switch m.manager.Type() {
	//case remote.IPC:
	//	return ipc.PythonExecutorTemplate, true
	//default:
	//	return "", false
	//}
	return "", true
}

func (m *VenvManager) createEnv(name string) (*VirtualEnv, error) {
	venvPath := path.Join(venvPath, name)

	if err := os.MkdirAll(venvPath, 0755); err != nil {
		return nil, errors.WrapWith(err, "venv %s: path creation failed", name)
	}

	cmd := exec.Command(pythonExec, "-m", "venv", venvPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, errors.WrapWith(err, "venv %s: venv creation failed", name)
	}

	if template, ok := m.template(); !ok {
		return nil, errors.Format("venv %s: no template is found for python", name)
	} else if err := os.WriteFile(path.Join(venvPath, venvStart), []byte(template), 0644); err != nil {
		return nil, errors.WrapWith(err, "venv %s: template creation failed", name)
	}

	return &VirtualEnv{
		ctx: m.ctx,
		//handler:  m.manager.NewExecutor(name),
		Name:     name,
		Exec:     path.Join(venvPath, "bin", "python"),
		Packages: []string{"pyzmq"},
	}, nil
}

func (m *VenvManager) addVenv(venv *VirtualEnv) {
	m.Envs[venv.Name] = venv
}

func (m *VenvManager) GetVenv(name string, requirements ...string) (*VirtualEnv, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if env, ok := m.Envs[name]; ok {
		//env.Run(m.Addr())
		return env, nil
	}

	env, err := m.createEnv(name)
	if err != nil {
		return nil, err
	}

	if err = env.AddPackages(requirements...); err != nil {
		return nil, err
	}

	m.addVenv(env)

	//env.Run(m.Addr())
	return env, nil
}
