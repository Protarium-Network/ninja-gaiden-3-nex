# Protocol coverage

| Protocol | Coverage |
|---|---|
| Ticket Granting | Login / LoginEx / RequestTicket (stock) |
| Secure Connection | stock (insecure Register enabled for Wii U) |
| Utility | stock (AcquireNexUniqueID …) |
| Ranking | UploadScore, UploadCommonData, GetRanking (Range / Near / User), GetCommonData, GetCachedTopX; Postgres `ng3_*` tables |
| DataStore | PreparePostObject / CompletePostObject / PrepareGetObject / GetMeta / search; ratings stub; objects in `datastore.objects`, blobs in S3 |
| NAT Traversal | stock |
| MatchMaking / MatchMakingExt / MatchmakeExtension | stock common handlers; MatchMaking lib pinned 3.3.0; pre-3.4 `AutoMatchmakeWithSearchCriteria_Postpone` decoder and friends-only `AutoMatchmakeWithGatheringID_Postpone` handler ported from the Sochi 2014 server |

Not registered: Notification, Monitoring, Health, AccountManagement
(housekeeping clients the game does not need answered).

Tested on real consoles: login, secure register, NAT traversal, room
creation/matchmaking, ranking and DataStore calls all went through. Not every
game mode has been exercised. The `[NG3 Secure]` log lines show which methods
a given mode calls.
