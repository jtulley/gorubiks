# gorubiks rubiks cube simulator

My son pointed out that, mathematically, any repeated pattern of rubiks cube moves should result in the cube getting back to the original state.  The first pattern I tried did not do so after 1000 moves, so I wrote this simulator to prove / disprove the theory, and to learn a bit more about writing in golang.

It turns out I just wasn't patient enough.  1260 was the right number.

Note: as of this writing, main.go does nothing, and everything is test-driven.  A better notation for initialization and rotation should be made, and a command line created around all of that.

