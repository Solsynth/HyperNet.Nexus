package web

import (
	"bufio"
	"os"

	"github.com/rs/zerolog/log"
)

var ipBlocklist []string

func AddBlockIP(ip string) {
	ipBlocklist = append(ipBlocklist, ip)
}

func ParseBlockIPList(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ipBlocklist = append(ipBlocklist, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	log.Info().Int("count", len(ipBlocklist)).Msg("Block IP list parsed successfully.")
	return nil
}
