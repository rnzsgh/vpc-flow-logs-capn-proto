package main

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	flowlogs "github.com/vpc-flow-logs-compression/vpc"

	"zombiezen.com/go/capnproto2"
)

func main() {

	writeCapnpFile()

	readCapnpFile()
}

func readCapnpFile() {

	file, err := os.Open("capnp.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gzIn, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer gzIn.Close()

	decoder := capnp.NewPackedDecoder(gzIn)

	counter := 0

	for {

		_, err := decoder.Decode()

		if err != nil {
			if !strings.Contains(err.Error(), "EOF") {
				panic(err)
			}
			break
		}

		counter++
	}

	fmt.Println(counter)
}

func writeCapnpFile() {

	file, err := os.Open("raw.txt.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	out, err := os.OpenFile("capnp.gz", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	gzOut, err := gzip.NewWriterLevel(out, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	gzOutBuffer := bufio.NewWriter(gzOut)

	gzIn, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer gzIn.Close()

	packedEncoder := capnp.NewPackedEncoder(gzOutBuffer)

	counter := 0

	scanner := bufio.NewScanner(gzIn)
	for scanner.Scan() {
		counter++

		values := strings.Split(scanner.Text(), " ")

		msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
		if err != nil {
			panic(err)
		}

		entry, err := flowlogs.NewRootVpcFlowLogEntry(seg)

		if err != nil {
			panic(err)
		}

		version, err := strconv.ParseInt(values[0], 10, 8)
		if err != nil {
			panic(err)
		}
		entry.SetVersion(int8(version))

		accountId, err := strconv.ParseUint(values[1], 10, 64)
		if err != nil {
			panic(err)
		}
		entry.SetAccountId(accountId)

		entry.SetInterfaceId(values[2])
		entry.SetSrcAddr(ipStringToInt(values[3]))
		entry.SetDstAddr(ipStringToInt(values[4]))

		srcPort, err := strconv.ParseUint(values[5], 10, 16)
		if err != nil {
			panic(err)
		}
		entry.SetSrcPort(uint16(srcPort))

		dstPort, err := strconv.ParseUint(values[6], 10, 16)
		if err != nil {
			panic(err)
		}
		entry.SetDstPort(uint16(dstPort))

		protocol, err := strconv.ParseUint(values[7], 10, 8)
		if err != nil {
			panic(err)
		}
		entry.SetProtocol(uint8(protocol))

		packets, err := strconv.ParseUint(values[8], 10, 16)
		if err != nil {
			panic(err)
		}
		entry.SetPackets(uint16(packets))

		bytes, err := strconv.ParseUint(values[9], 10, 64)
		if err != nil {
			panic(err)
		}
		entry.SetBytes(bytes)

		start, err := strconv.ParseUint(values[10], 10, 32)
		if err != nil {
			panic(err)
		}
		entry.SetStart(uint32(start))

		end, err := strconv.ParseUint(values[11], 10, 32)
		if err != nil {
			panic(err)
		}
		entry.SetEnd(uint32(end))

		if values[12] == "ACCEPT" {
			entry.SetAction(flowlogs.VpcFlowLogEntry_Action_accept)
		} else {
			entry.SetAction(flowlogs.VpcFlowLogEntry_Action_reject)
		}

		if values[13] == "OK" {
			entry.SetLogStatus(flowlogs.VpcFlowLogEntry_LogStatus_ok)
		} else if values[13] == "NODATA" {
			entry.SetLogStatus(flowlogs.VpcFlowLogEntry_LogStatus_noData)
		} else {
			entry.SetLogStatus(flowlogs.VpcFlowLogEntry_LogStatus_skipData)
		}

		err = packedEncoder.Encode(msg)
		if err != nil {
			panic(err)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(counter)

	gzOutBuffer.Flush()

	gzOut.Flush()

}

func ipStringToInt(str string) uint32 {

	ip := net.ParseIP(str)

	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}
