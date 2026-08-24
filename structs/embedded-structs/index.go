// example
// type car struct {
//   brand string
//   model string
// }

// type truck struct {
//   // "car" is embedded, so the definition of a
//   // "truck" now also additionally contains all
//   // of the fields of the car struct
//   car
//   bedSize int
// }

// nested vs embedded:-
// nested : tuck.car.brand
// embeded: truck.brand

package main

type sender struct {
	rateLimit int
	user      // this struct is embedded; now the properties or fields of user can be directly accesse (composition)
}

type user struct {
	name   string
	number int
}
