# ninja-gaiden-3-nex

A NEX server for **Ninja Gaiden 3: Razor's Edge** on Wii U. The official
servers are gone, so this brings back what the game used them for: the
leaderboards, replay/clip sharing and online rooms (co-op, Ninja Trials,
clan).

It's built on the Pretendo NEX libraries (`nex-go`, `nex-protocols-go`,
`nex-protocols-common-go`) and follows the same layout as the other Wii U
servers in this org.

The game has been played online on real consoles against this server. It
hasn't been tested exhaustively, so expect rough edges in the less common
modes.

## How the game talks to the server

Nothing exotic. The executable statically links the NEX client and uses
these protocols, all of which are registered here:

- Ticket Granting + Secure Connection (login, secure handoff)
- NAT Traversal
- MatchMaking, MatchMaking Ext, Matchmake Extension (rooms)
- Ranking (leaderboards)
- DataStore (replays/clips; blobs go to an S3-compatible bucket)
- Utility

Where each constant came from is written down in [RECON.md](RECON.md), and
the `recon/` folder has the scripts used to read them out of the RPX.

## Configuration

| | |
|---|---|
| Access key | `f857b4bd` |
| NEX version | 3.4.13 |
| MatchMaking library | pinned to 3.3.0 (this game uses the pre-3.4 session layout) |
| Ranking library | bumped to 3.6.0 so the update time field is serialized |
| Structure headers | off |
| Legacy connection signature | **on** (without it the console shows 106-0502) |

### Game server IDs

One per title ID. The PRUDP side doesn't care which one a console came from,
but your token/NASC front-end has to map **all of them** to this server or
only some regions will connect.

| Region | Title ID | Game server ID |
|---|---|---|
| JPN | `0005000010110900` | `10110900` (read from the executable) |
| USA | `0005000010110A00` | `10110a00` |
| EUR | `0005000010110B00` | `10110b00` |
| EUR (second release) | `0005000010139B00` | `10139b00` |

The USA and EUR values are the unique-ID half of the title ID, which is what
the JPN build does. I didn't have those executables to confirm it.

## Running it

Docker is the easy way:

```bash
cp .env.example .env
cp settings.example.json settings.json
docker compose up --build
```

That starts Postgres and the server on UDP 28800 (auth) and 28801 (secure).
In the default local mode there's no account server: put your console's PID
and NEX password in `settings.json` and the login token isn't checked. Don't
expose that to the internet.

For a shared deployment set `PN_NG3_LOCAL_MODE=0` and give it
`PN_NG3_NEX_TOKEN_AES_KEY` (64 hex chars) and `PN_NG3_NEX_PASSWORD_SECRET`
(at least 32 bytes, hex). Both must match whatever issues your NEX tokens.
Set `PN_NG3_SECURE_HOST` to the address consoles should use to reach the
secure port; keep it short, the game truncates it.

Replay uploads need an S3-compatible bucket (`PN_S3_ENDPOINT`,
`PN_S3_ACCESS_KEY`, `PN_S3_SECRET_KEY`, `PN_S3_BUCKET`). MinIO is fine.
Without it everything else works, uploads fail.

Without Docker: `go build -o ng3-nex .` and export the variables from
`compose.yaml`, including `PN_NG3_POSTGRES_URI`.

The matchmaking tables are created by the library on first start; the
ranking and DataStore tables are created by `database/init_postgres.go`.

## Notes

- `internal/nex-protocols-common-go-patch/` is a vendored copy of
  `nex-protocols-common-go` with the DataStore S3 presigner and a few
  matchmaking fixes. It's wired in through a `replace` in `go.mod`.
- The pre-3.4 `AutoMatchmakeWithSearchCriteria_Postpone` decoder and the
  friends-only `AutoMatchmakeWithGatheringID_Postpone` handler in `nex/` are
  carried over from the Mario & Sonic Sochi 2014 server, same NEX build.
- Every RMC call is printed (`[NG3 Secure] PID=… protocol=… method=…`),
  which is how you find out what a game mode needs if something isn't
  answered.
- Ranking leaderboards for this game aren't friend-aware; friend-range
  queries fall back to the user query.

## License

AGPL-3.0, same as the Pretendo libraries it's built on. See
[LICENSE](LICENSE) and [NOTICE](NOTICE). No Nintendo or Tecmo Koei code or
assets are included.
