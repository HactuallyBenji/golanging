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

func TestSubtractFormatAmount(t *testing.T) {
   ans := SubtractFormatAmount(4.00, 2.00)
   if ans != "USD 2.00" {

      t.Errorf("FormatAmount(4.00, 2.00) = %s; Should be USD 2.00", ans)
   }
}

func TestSubtractFormatAmount2(t *testing.T) {
   ans := SubtractFormatAmount(3.00, 1.12)
   if ans != "USD 1.88" {
      t.Errorf("FormatAmount(3.00,1.12) = %s; Should be USD 1.88", ans) // we use t to record the testing error.
   }
}

func TestSubtractFormatAmount3(t *testing.T) {
   ans := SubtractFormatAmount(1.00, 1.12)
   if ans != "Impossible operation" {
      t.Errorf("FormatAmount(1.00, 1.12) cannot be performed")
   }
}
