package main

import (
   "testing"
)
func TestFormatAmount(t *testing.T) {
   ans := FormatAmount(2.00)
   if ans != "USD 2.00" {
     // we use t to record the testing error
     t.Errorf("FormatAmount(2.00) = %s; Should be 2.00", ans)    }
}
