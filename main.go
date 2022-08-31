package main

import (
	"machine"
	"time"
)

var (
	uart = machine.UART0
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
)

func main() {
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})
	for {
		uart.WriteByte('A')
		time.Sleep(time.Millisecond * 5000)
		uart.WriteByte('B')
		time.Sleep(time.Millisecond * 5000)
	}
}
