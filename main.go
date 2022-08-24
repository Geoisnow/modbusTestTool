package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/goburrow/modbus"
)

type Slave struct {
	Address      string
	SlaveId      int
	BaudRate     int
	Parity       string
	DataBits     int
	StopBits     int
	Timeout      int
	IdleTimeout  int
	RegisterAddr int
	Quantity     int
}

var slave Slave

func init() {
	flag.StringVar(&slave.Address, "a", "/dev/ttyS0", "Serial address")
	flag.IntVar(&slave.SlaveId, "i", 1, "Slave id")
	flag.IntVar(&slave.BaudRate, "b", 9600, "BaudRate")
	flag.IntVar(&slave.DataBits, "d", 8, "DataBits ")
	flag.IntVar(&slave.StopBits, "s", 1, "StopBits")
	flag.StringVar(&slave.Parity, "p", "N", "Parity")
	flag.IntVar(&slave.Timeout, "t", 100, "Timeout")
	flag.IntVar(&slave.IdleTimeout, "it", 2, "IdleTimeout")
	//读取寄存器相关
	flag.IntVar(&slave.RegisterAddr, "r", 0, "Register addr (default 0)")
	flag.IntVar(&slave.Quantity, "q", 1, "Quantity")
}

func main() {
	flag.Parse()
	handler := NewRtuHandler()
	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		log.Println("Handler connect err:", err)
		fmt.Printf("Addr:%s\tSlaveId:%d\tBaudRate:%d\nParity:%s\tDataBits:%d\tStopBits:%d\nTimeout:%dms\tIdleTimeout：%ds\n",
			handler.Address, handler.SlaveId, handler.BaudRate, handler.Parity, handler.DataBits, handler.StopBits, handler.Timeout/1000, handler.IdleTimeout/1000000000)
		return
	}

	client := modbus.NewClient(handler)
	res, err := client.ReadHoldingRegisters(uint16(slave.RegisterAddr), uint16(slave.Quantity))
	if err != nil {
		log.Println("ReadHoldingRegisters err:", err)
		fmt.Printf("RegisterAddr:%d\tQuantity:%d\n", slave.RegisterAddr, slave.Quantity)
		return
	}
	fmt.Println("Read success!")
	fmt.Println("Result :", res)
	return
}

func NewRtuHandler() *modbus.RTUClientHandler {
	handler := modbus.NewRTUClientHandler(slave.Address)
	handler.SlaveId = byte(slave.SlaveId)
	handler.BaudRate = slave.BaudRate
	handler.Parity = slave.Parity
	handler.DataBits = slave.DataBits
	handler.StopBits = slave.StopBits
	handler.Timeout = time.Duration(slave.Timeout) * time.Millisecond
	handler.IdleTimeout = time.Duration(slave.IdleTimeout) * time.Second
	return handler
}
