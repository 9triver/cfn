package platform

import (
	"context"
	"fmt"
	messages "github.com/9triver/ignis/proto/controller"
	"github.com/asynkron/protoactor-go/actor"
	"log/slog"
)

type Controller struct {
	name       string
	pid        *actor.PID
	behavior   *actor.Behavior
	envManager *VenvManager
}

func NewController() *Controller {
	manager, err := NewManager(context.Background())
	if err != nil {
		return nil
	}
	controller := &Controller{
		name:       "work-platform/controller",
		pid:        nil,
		behavior:   &actor.Behavior{},
		envManager: manager,
	}
	controller.behavior.Become(controller.initBehavior)
	return controller
}

func (controller *Controller) Receive(context actor.Context) {
	logger := context.Logger()
	logger.Info("Receive")
	controller.behavior.Receive(context)
}

//======================================================================================================================
// Behavior
//----------------------------------------------------------------------------------------------------------------------

func (controller *Controller) initBehavior(context actor.Context) {
	logger := context.Logger()
	switch msg := context.Message().(type) {
	case *actor.Started:
		controller.handleInit(context) // 初始化
	default:
		controller.handleUnexpectedMessage(msg, logger) // TODO: 返回提示给发送者
	}
}

func (controller *Controller) runningBehavior(context actor.Context) {
	logger := context.Logger()
	logger.Info("Running behavior")
	switch msg := context.Message().(type) {
	case *messages.AppendPyFunc:
		controller.handleAppendPyFunc(msg, logger)
	default:
		controller.handleUnexpectedMessage(msg, logger)
	}
}

//======================================================================================================================
// Handler
//----------------------------------------------------------------------------------------------------------------------

func (controller *Controller) handleAppendPyFunc(msg *messages.AppendPyFunc, logger *slog.Logger) {
	// 初始化虚拟环境
	venv, err := controller.envManager.GetVenv(msg.GetName(), msg.GetRequirements()...)
	if err != nil {
		logger.Error(fmt.Sprint(err))
		return
	}
	logger.Info("虚拟环境初始化完成\n\t" + fmt.Sprint(venv))
}

func (controller *Controller) handleInit(context actor.Context) {
	controller.pid = context.Self()
	controller.behavior.Become(controller.runningBehavior)
}

func (controller *Controller) handleUnexpectedMessage(msg interface{}, logger *slog.Logger) {
	logger.Warn(fmt.Sprintf("%v: Unexpected message type received. %v", controller.name, msg))
}
