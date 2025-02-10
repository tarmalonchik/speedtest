package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	iperf3client "github.com/tarmalonchik/speedtest/internal/app/unit/workers/iperf3-client"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
)

func main() {
	ctx := context.Background()

	err := measureSingleNode(ctx, "95.217.134.47")
	fmt.Println(err)
}

func measureSingleNode(ctx context.Context, ip string) error {
	const divider = "- - - - - - - - - - - - - - - - - - - - - - - - -"
	const finisher = "iperf Done."

	data, err := exec.CommandContext(ctx, "iperf3", "-c", ip, "-p", "5201", "-t5", "--json").Output()
	if err != nil {
		return trace.FuncNameWithErrorMsg(err, "executing command")
	}

	var payload iperf3client.IperfJsonOut

	if err = json.Unmarshal(data, &payload); err != nil {
		return trace.FuncNameWithErrorMsg(err, "unmarshal")
	}
	fmt.Println(payload.End.SumReceived.BitsPerSecond)
	fmt.Println(payload.End.SumSent.BitsPerSecond)
	return nil
}
