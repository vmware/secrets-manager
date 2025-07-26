/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"bufio"
	"os"
)

func (i *Initializer) commandFileScanner(
	cid *string) (*os.File, *bufio.Scanner) {
	filePath := i.EnvReader.InitCommandPathForSentinel()
	file, err := i.FileOpener.Open(filePath)

	if err != nil {
		i.Logger.InfoLn(
			cid,
			"RunInitCommands: no initialization file found... "+
				"skipping custom initialization.",
		)
		return nil, nil
	}

	i.Logger.TraceLn(cid, "Before parsing commands 001")

	return file, bufio.NewScanner(file)
}
