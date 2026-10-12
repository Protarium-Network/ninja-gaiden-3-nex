# Ninja Gaiden 3: Razor's Edge (Wii U) — NEX server

Preservation NEX server for **Ninja Gaiden 3: Razor's Edge** (Team Ninja /
Tecmo Koei, 2012) on Wii U, so its online features — **leaderboards** (story,
chapter challenge, Ninja Trials, clan), **replay/clip sharing** and
**online co-op / clan rooms** — can work again. Built on the
[Pretendo Network](https://github.com/PretendoNetwork) NEX libraries
(`nex-go`, `nex-protocols-go`, `nex-protocols-common-go`), same layout as the
other Protarium Wii U servers.

## Status

Builds and passes its unit test. **Not yet verified on a console.** Every
constant was recovered from the retail executable ([RECON.md](RECON.md));
protocol status is in [PROTOCOL_COVERAGE.md](PROTOCOL_COVERAGE.md).

## Recovered configuration

| Field | Value |
|---|---|
| Access key | `f857b4bd` |
| NEX version | `3.4.13` (`ngs/3.4.x.3`) |
| MatchMaking lib | pinned to `3.3.0` (pre-3.4 session layout) |
| Ranking lib | bumped to `3.6.0` (so `RankingRankData.UpdateTime` serializes) |
| `UseStructureHeader` | `false` |
| `LegacyConnectionSignature` | `false` by default — set `PN_NG3_LEGACY_SIGNATURE=1` if the console shows 106-0502 |

### Game server IDs (all regions)

| Region | Title ID | Game server ID |
|---|---|---|
| JPN | `0005000010110900` | `10110900` (read from the binary) |
| USA | `0005000010110A00` | `10110a00` (derived) |
| EUR | `0005000010110B00` | `10110b00` (derived) |
| EUR (2nd release) | `0005000010139B00` | `10139b00` (derived) |

The PRUDP server is region-agnostic (one shared key, same ports for all).
Register **every** ID in your NASC / account-server front-end, all pointing
at the same auth port — otherwise only one region works.
`globals.GameServerIDs` lists them for tooling.

## Running

```bash
cp .env.example .env
cp settings.example.json settings.json   # local mode: your console's PID + NEX password
docker compose up --build                # UDP 28800 (auth) / 28801 (secure)
```

Without Docker: `go build -o ng3-nex . && ./ng3-nex` with the `PN_NG3_*`
variables from `compose.yaml` and a PostgreSQL URI in `PN_NG3_POSTGRES_URI`.

Shared mode: `PN_NG3_LOCAL_MODE=0` plus `PN_NG3_NEX_TOKEN_AES_KEY` (64 hex)
and `PN_NG3_NEX_PASSWORD_SECRET` (≥32 bytes hex) matching your account
server. DataStore uploads (replays/clips) need an S3-compatible bucket
(`PN_S3_*`; MinIO works).

`internal/nex-protocols-common-go-patch/` is a vendored fork carrying the
DataStore/S3 and matchmaking changes this stack needs.

## License

AGPL-3.0, see [LICENSE](LICENSE) and [NOTICE](NOTICE). No proprietary code or
assets are included; `recon/` only holds the analysis scripts.
