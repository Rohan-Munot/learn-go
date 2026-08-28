// package main

// type contact struct {
// 	sendingLimit int32
// 	userID       string
// 	age          int32
// }

// type perms struct {
// 	canSend         bool
// 	canReceive      bool
// 	permissionLevel int
// 	canManage       bool
// }

package main

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	canSend         bool
	canReceive      bool
	canManage       bool
	permissionLevel int
}
