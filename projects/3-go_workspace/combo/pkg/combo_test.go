package pkg

import "testing"

func TestUse(t *testing.T) {
    if err := Use(); err != nil {
        t.Fatal(err)
    }
}
