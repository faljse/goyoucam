package main

import (
	"bytes"
	"log"
	"os/exec"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type CmdWatcher struct {
	RunningCmd []RunningCmd
	hub *Hub
}

type CaptureMpeg struct {
	hub *Hub

}

func newCaptureMpeg(hub * Hub) *CaptureMpeg {
	return &CaptureMpeg{
		hub:hub,
	}
}

func (capt *CaptureMpeg) Write(p []byte) (n int, err error) {
	capt.hub.write(p)
	return len(p),nil
}

type RunningCmd struct {
	Cmd  Cmd
	WBuf []byte
	lines []string
	hub *Hub
	cMpeg *CaptureMpeg
}

func newRunningCmd(hub *Hub) *RunningCmd{
	return &RunningCmd{
		hub:hub,
		cMpeg: newCaptureMpeg(hub),
	}
}

func (rcmd *RunningCmd) Write(p []byte) (n int, err error) {
	rcmd.WBuf=append(rcmd.WBuf, p...)
	for {
		idx := bytes.IndexRune(rcmd.WBuf, '\n')
		if idx == -1 {
			return len(p), nil
		}
		line := rcmd.WBuf[0:idx]
		rcmd.WBuf = rcmd.WBuf[idx+1:]
		rcmd.lines = append(rcmd.lines, string(line))
		log.Println(string(line))
	}
}

func (cw *CmdWatcher) addComd(c Cmd) {
	cmd := exec.Command(c.Cmd, c.Args...)
	rcmd := newRunningCmd(cw.hub)
	cmd.Stdout = rcmd.cMpeg
	cmd.Stderr = rcmd
	errCmd := cmd.Start()
	log.Printf("Command %v finished with error: %v", c.Cmd, errCmd)
}

func newCmdWatcher(hub *Hub) *CmdWatcher {
	return &CmdWatcher{
		hub:hub,
	}
}
