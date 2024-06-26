// SPDX-FileCopyrightText: Copyright (c) 2023-2024, CIQ, Inc. All rights reserved
// SPDX-License-Identifier: Apache-2.0

package utils

import "time"

const timeFormat = "2006-01-02 15:04:05 MST"

func TimeToString(t int64) string {
	if t == 0 {
		return ""
	}
	return time.Unix(t, 0).Format(timeFormat)
}

func StringToTime(s string) int64 {
	if s == "" {
		return 0
	}
	t, _ := time.Parse(timeFormat, s)
	return t.Unix()
}
