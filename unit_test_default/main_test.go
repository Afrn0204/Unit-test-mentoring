package main

import "testing"

func TestLuasBelahKetupat(t *testing.T) {
	diagonal1 := 10.0
	diagonal2 := 8.0
	expected := 40.0

	result := luasBelahKetupat(diagonal1, diagonal2)

	if result != expected {
		t.Errorf("Hasil salah, dapat: %.2f, seharusnya: %.2f", result, expected)
	}

	diagonal1 = 12.0
	diagonal2 = 6.0
	expected = 36.0

	result = luasBelahKetupat(diagonal1, diagonal2)

	if result != expected {
		t.Errorf("Hasil salah, dapat: %.2f, seharusnya: %.2f", result, expected)
	}
}
