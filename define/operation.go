package define

import "comet/libs"

const (
	OP_AUTH = 1
	OP_AUTH_REPLY = 2

	OP_JOIN_ROOM = 3
	OP_JOIN_ROOM_REPLY = 4

	OP_EXIT_ROOM = 5
	OP_EXIT_ROOM_REPLY = 6
)

type BoardcastRoomArg struct {
	Rid string
	P libs.Proto
}