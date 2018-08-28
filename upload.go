/*
 * Copyright (c) 2018. The speakin Authors. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */
package voiceprint_cloud_go_git

type UploadFileReq struct {
	Type       string `json:"Type" binding:"required"`
	DurationMS int64  `json:"durationMS"`
	Timestamp  int64  `json:"timestamp" binding:"required"`
}
