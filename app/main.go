/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

/*
#include <stdlib.h>

int sum(int a, int b) {
    return a + b;
}
*/
import "C"
import "fmt"

func main() {
	a, b := 3, 5
	result := C.sum(C.int(a), C.int(b))
	fmt.Printf("%d + %d = %d\n", a, b, result)
}
