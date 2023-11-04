package main

import (
   "testing"
)
func TestFormatAmount(t *testing.T) {
   ans := FormatAmount(2.00)
   if ans != "USD 2.00" {
     // we use t to record the testing error
     t.Errorf("FormatAmount(2.00) = %s; Should be USD 2.00", ans)    }
}

func TestFormatAmount2(t *testing.T) {
   ans := FormatAmount(4.00)
   if ans != "USD 4.00" {
      t.Errorf("FormatAmount(4.00) = %s; Should be USD 4.00", ans)
   }
}

func TestFormatAmount3(t *testing.T) {
   ans := FormatAmount(5.10)
   if ans != "USD 5.10" {
      t.Errorf("FormatAmount(5.10) = %s; Should be USD 5.10", ans)
   }
}
