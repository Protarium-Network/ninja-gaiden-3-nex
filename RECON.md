# Recon — Ninja Gaiden 3: Razor's Edge (Wii U)

Source: `sigma3_cafe.rpx` (JPN build, game server ID `0x10110900`; Team Ninja
internal name "sigma3"). Everything below was read from the executable.
Reproduce with the scripts in `recon/` (see `recon/README.md`).

## Constants

| Field | Value | Evidence |
|---|---|---|
| Access key | `f857b4bd` | UTF-16BE literal at `.rodata 0x1002e9e4` (single occurrence), loaded by the NEX init function at `0x02312130` (`lis r6,0x1003 ; addi r6,r6,-0x161c`) and stored in the init-parameter struct. |
| Game server ID | `0x10110900` | `lis r5,0x1011 ; addi r5,r5,0x900` at `0x02312128`, stored next to the key in the same struct (`0x30(r1)` / `0x38(r1)`). Equals the unique-ID half of the JPN title `0005000010110900`. |
| NEX build | `3.4.13` (`ngs/3.4.x.3`) | kinnay.github.io `nexwiiu.json`: id 269551872, key `f857b4bd`, `build:3_4_13_3_0`. The `NEX_3_0_1_5` / `NEX_RK_3_0_1` / `NEX_MM_3_0_1` strings in the binary are per-protocol library tags, not the NGS version. |
| Library tags | `NEX_3_0_1_5`, `NEX_DS_3_0_1`, `NEX_RK_3_0_1`, `NEX_MM_3_0_1`, `NEX_UT_3_0_1` | `.rodata`: base NEX, DataStore, Ranking, MatchMaking, Utility. |

### Regions

| Region | Product code | Title ID | Game server ID | Source |
|---|---|---|---|---|
| JPN | WUP-P-ANGJ | `0005000010110900` | `10110900` | read from the binary; wiiubrew title DB |
| USA | WUP-P-ANGE | `0005000010110A00` | `10110a00` | wiiubrew title DB (derived: unique-ID half) |
| EUR | WUP-P-ANGP | `0005000010110B00` | `10110b00` | wiiubrew title DB (derived) |
| EUR (2nd) | WUP-P-ANGZ | `0005000010139B00` | `10139b00` | wiiubrew title DB (derived) |

The binary's ID equals its own title-ID half, so the same rule is assumed for
the other regions. **USA/EUR IDs stay unconfirmed until a US/EU console or
dump is checked** (DKC:TF is a counter-example where one ID covers all
regions, so registering all four is the safe choice). Only the nex_token
front-end depends on them.

## What the game links

`nn::nex` client classes present: `Authentication`/TicketGranting,
`SecureConnection`, `AccountManagement`, `Health`, `Monitoring`,
`Notification`, `NintendoNotificationEvent`, `RemoteLogDevice`,
`NATTraversal` (+Relay, ReportInternal), `MatchMakingProtocol`,
`MatchMakingProtocolExt`, `MatchmakeExtension`, `Ranking`, `Utility`,
`DataStore` (+Hpp/LogicServer/FileServer helper clients).

Game-side class names (RTTI strings) show what each is for:

- **Ranking** — `Online::RankingMission/Total/Trials/Clan/ChapterChallenge`,
  `RankingReadProc`, `RankingRegisterStoryRankProc`,
  `RankingRegisterChapterChallengeRankProc`, `RankingReadRankFisrtScoreProc`,
  `TFW::Network::Ranking::RequestOfDownloadRangeByRank` / `…GameData`.
- **Replays / clips** — `Online::RankingReplayUploadProc`,
  `RankingClipReadProc` (DataStore objects), `OnlineUserStorage`.
- **Matchmaking** — `Online::OnlineMatching`, `RoomCreate` / `RoomJoin`
  (+`Clan`, `Trials`), `OnlineInvitation`, `OnlineFriend`, co-op / clan
  battle packets (`NotifyIngameStartPacket`, `NotifyTrialsFinishPacket`, …)
  carried peer-to-peer after NAT traversal.
- **Voice chat** — `OnlineVoicechat` (peer-to-peer, no server involvement).

## Server consequences

Auth + Secure endpoints; secure registers SecureConnection, Utility, NAT
Traversal, MatchMaking, MatchMakingExt, MatchmakeExtension, Ranking,
DataStore. The executable is stripped, so which exact RMC methods the game
calls was not mapped statically; every call is logged
(`[NG3 Secure] PID=… protocol=… method=…`) instead.

`LegacyConnectionSignature` has to be `true` for this title: with `false` the
console never accepts the CONNECT-ACK and reports 106-0502.
