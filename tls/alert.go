// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import "strconv"

type alert uint8

const (
	// alert level
	alertLevelWarning = 1
	alertLevelError   = 2
)

const (
	alertCloseNotify            alert = 0
	alertUnexpectedMessage      alert = 10
	alertBadRecordMAC           alert = 20
	alertDecryptionFailed       alert = 21
	alertRecordOverflow         alert = 22
	alertDecompressionFailure   alert = 30
	alertHandshakeFailure       alert = 40
	alertBadCertificate         alert = 42
	alertUnsupportedCertificate alert = 43
	alertCertificateRevoked     alert = 44
	alertCertificateExpired     alert = 45
	alertCertificateUnknown     alert = 46
	alertIllegalParameter       alert = 47
	alertUnknownCA              alert = 48
	alertAccessDenied           alert = 49
	alertDecodeError            alert = 50
	alertDecryptError           alert = 51
	alertProtocolVersion        alert = 70
	alertInsufficientSecurity   alert = 71
	alertInternalError          alert = 80
	alertInappropriateFallback  alert = 86
	alertUserCanceled           alert = 90
	alertNoRenegotiation        alert = 100
	alertMissingExtension       alert = 109
	alertUnsupportedExtension   alert = 110
	alertUnrecognizedName       alert = 112
	alertNoApplicationProtocol  alert = 120

	// specific types for GMT0024
	alertUnspporttedSite2Site alert = 200
	alertNoArea               alert = 201
	alertUnspportedAreaType   alert = 202
	alertBadIBCParam          alert = 203
	alertUnspportedIBCParam   alert = 204
	alertIdentityNeed         alert = 205
)

var alertText = map[alert]string{
	alertCloseNotify:            "close notify",
	alertUnexpectedMessage:      "unexpected message",
	alertBadRecordMAC:           "bad record MAC",
	alertDecryptionFailed:       "decryption failed",
	alertRecordOverflow:         "record overflow",
	alertDecompressionFailure:   "decompression failure",
	alertHandshakeFailure:       "handshake failure",
	alertBadCertificate:         "bad certificate",
	alertUnsupportedCertificate: "unsupported certificate",
	alertCertificateRevoked:     "revoked certificate",
	alertCertificateExpired:     "expired certificate",
	alertCertificateUnknown:     "unknown certificate",
	alertIllegalParameter:       "illegal parameter",
	alertUnknownCA:              "unknown certificate authority",
	alertAccessDenied:           "access denied",
	alertDecodeError:            "error decoding message",
	alertDecryptError:           "error decrypting message",
	alertProtocolVersion:        "protocol version not supported",
	alertInsufficientSecurity:   "insufficient security level",
	alertInternalError:          "internal error",
	alertInappropriateFallback:  "inappropriate fallback",
	alertUserCanceled:           "user canceled",
	alertNoRenegotiation:        "no renegotiation",
	alertMissingExtension:       "missing extension",
	alertUnsupportedExtension:   "unsupported extension",
	alertUnrecognizedName:       "unrecognized name",
	alertNoApplicationProtocol:  "no application protocol",

	//specific types for GM
	alertUnspporttedSite2Site: "不支持site2site",
	alertNoArea:               "没有保护域",
	alertUnspportedAreaType:   "不支持的保护域类型",
	alertBadIBCParam:          "接收到一个无效的ibc公共参数",
	alertUnspportedIBCParam:   "不支持ibc参数中定义的信息",
	alertIdentityNeed:         "缺少对方的ibc标识",
}

func (e alert) String() string {
	s, ok := alertText[e]
	if ok {
		return "tls: " + s
	}
	return "tls: alert(" + strconv.Itoa(int(e)) + ")"
}

func (e alert) Error() string {
	return e.String()
}
