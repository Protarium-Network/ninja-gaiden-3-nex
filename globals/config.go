package globals

import (
	"os"
	"strings"
)

// NEX configuration for "Ninja Gaiden 3: Razor's Edge" (Wii U, Team Ninja /
// Tecmo Koei, 2012).
//
// AccessKey, game server ID and NGS build come from the retail executable
// (sigma3_cafe.rpx, see RECON.md) and are cross-checked against
// kinnay.github.io's public Wii U NEX database (id 269551872, key f857b4bd,
// branch origin/release/ngs/3.4.x.3, build 3_4_13_3_0).
//
// The game ships under several title IDs; each region's game server ID is the
// unique-ID half of its title ID. The PRUDP server is region-agnostic (one
// shared access key); the IDs only matter to the NASC / nex_token front-end,
// which must map EVERY ID below to this server's auth port.
const (
	AccessKey = "f857b4bd"

	GameServerIDJPN  = "10110900" // title 0005000010110900 (WUP-P-ANGJ) - in the binary
	GameServerIDUSA  = "10110a00" // title 0005000010110A00 (WUP-P-ANGE)
	GameServerIDEUR  = "10110b00" // title 0005000010110B00 (WUP-P-ANGP)
	GameServerIDEUR2 = "10139b00" // title 0005000010139B00 (WUP-P-ANGZ, 2nd EUR release)
)

// GameServerIDs lists every region this server is meant to serve.
var GameServerIDs = []string{GameServerIDJPN, GameServerIDUSA, GameServerIDEUR, GameServerIDEUR2}

// kinnay build tag "build:3_4_13_3_0" (origin/release/ngs/3.4.x.3), same NGS
// generation as Trine 2 / Sochi 2014 / DKC:TF.
//
// MatchMaking is pinned to 3.3.0 (pre-3.4 MatchmakeSession layout: no
// ProgressScore byte, as console-captured on Trine 2 and Sochi 2014). Ranking
// is bumped to 3.6.0 in nex/secure.go so RankingRankData.UpdateTime is
// serialized.
const (
	NEXMajor = 3
	NEXMinor = 4
	NEXPatch = 13

	MatchMakingMajor = 3
	MatchMakingMinor = 3
	MatchMakingPatch = 0
)

// LegacyConnectionSignature selects the PRUDPv1 CONNECT-ACK signature scheme.
// This title needs true: with false the console gets no valid CONNECT-ACK and
// fails with 106-0502. (Other 3.4.13 titles in this family behave the other
// way around, so don't copy this value blindly.) PN_NG3_LEGACY_SIGNATURE=0|1
// overrides it.
var LegacyConnectionSignature = true

func init() {
	if v := strings.TrimSpace(os.Getenv("PN_NG3_LEGACY_SIGNATURE")); v != "" {
		LegacyConnectionSignature = v == "1"
	}
}
