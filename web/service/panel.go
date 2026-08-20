package service

import (
	"os"
	"syscall"
	"time"

	"github.com/mhsanaei/3x-ui/v2/logger"
)

// PanelService provides system-level panel controls.
type PanelService struct{}

// RestartPanel schedules a graceful panel restart after the current request returns.
func (s *PanelService) RestartPanel(delay time.Duration) error {
	p, err := os.FindProcess(syscall.Getpid())
	if err != nil {
		return err
	}
	go func() {
		time.Sleep(delay)
		if err := p.Signal(syscall.SIGHUP); err != nil {
			logger.Error("failed to send SIGHUP signal:", err)
		}
	}()
	return nil
}
