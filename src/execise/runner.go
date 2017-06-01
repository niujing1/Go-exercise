package main

//runner包管理处理任务的运行和生命周期

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//runner在给定的超时时间内执行一组任务
//并且在操作系统给出中断信号时结束这些任务
type Runner struct {
	// interrupt 通道包裹从操作系统发送的信号
	interrupt chan os.Signal

	//complete通道报告任务已处理完成
	complete chan error

	//timeout报告处理任务已超时
	timeout <-chan time.Time

	//函数
	tasks []func(int)
}

//ErrorTimeout会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

//ErrInterrupt会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

//New返回一个新准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add将任务附加到一个Runner上，这个任务是一个接受欧int型ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//start执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt) //接收所有中断信号

	//用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	//当任务处理完成时发出信号
	case err := <-r.complete:
		return err

		//当任务处理超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

//run执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测系统中的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		//执行以注册的任务
		task(id)
	}

	return nil
}

//gotInterrupt验证是否收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt) //停止接收后续的任何信号
		return true

		//否则继续执行
	default:
		return false
	}
}
