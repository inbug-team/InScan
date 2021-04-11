/*
负责人员：InBug Team
创建时间：2021/3/4
程序用途：
*/
package icmp

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func CheckAlive(host string) bool {
	size := 32
	var seq int16 = 1
	const EchoRequestHeadLen = 8
	const EchoReplyHeadLen = 20

	startTime := time.Now()
	conn, err := net.DialTimeout("ip4:icmp", host, 6*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	id0, id1 := genIdentifier(host)

	msg := make([]byte, size+EchoRequestHeadLen)
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4], msg[5] = id0, id1
	msg[6], msg[7] = genSequence(seq)

	length := size + EchoRequestHeadLen

	check := checkSum(msg[0:length])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	if err := conn.SetDeadline(startTime.Add(6 * time.Second)); err != nil {
		return false
	}

	if _, err := conn.Write(msg[0:length]); err != nil {
		return false
	}

	receive := make([]byte, EchoReplyHeadLen+length)
	if _, err := conn.Read(receive); err != nil {
		return false
	}

	return true
}

func checkSum(msg []byte) uint16 {
	sum := 0
	length := len(msg)
	for i := 0; i < length-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if length%2 == 1 {
		sum += int(msg[length-1]) * 256
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum = sum + (sum >> 16)
	answer := uint16(^sum)
	return answer
}

func genSequence(v int16) (byte, byte) {
	ret1 := byte(v >> 8)
	ret2 := byte(v & 255)
	return ret1, ret2
}

func genIdentifier(host string) (byte, byte) {
	return host[0], host[1]
}

func MultiCheckAlive(workerPing int, hosts []string) []string {
	n := workerPing
	if len(hosts) > workerPing {
		n = workerPing
	} else {
		n = len(hosts)
	}
	worker := make(chan bool, n)
	var wg sync.WaitGroup
	var lock sync.Mutex
	var hostsAlive []string
	for _, host := range hosts {
		wg.Add(1)
		worker <- true
		go func(_host string, _wg *sync.WaitGroup) {
			status := CheckAlive(_host)
			if status {
				lock.Lock()
				fmt.Println(_host, "is alive !")
				hostsAlive = append(hostsAlive, _host)
				lock.Unlock()
			}
			<-worker
			wg.Done()
		}(host, &wg)
	}
	wg.Wait()
	close(worker)
	return hostsAlive
}
