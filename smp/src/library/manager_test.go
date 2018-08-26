package library

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
		return
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}

	m0 := &MusicEntry{
		"1", "My Heart Will Go On", "Celion Dion",
		"http://qbox.me/xxxxxxxx", "MP3",
	}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist || m.Name != m0.Name ||
		m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}

	m, err := mm.Get(0)
	if err != nil {
		t.Error("NewMusicManager.Get() failed.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.", err)
	}

	m1 := &MusicEntry{
		"2", "Hello world1", "Celion Dion",
		"http://qbox.me/xxxxxxxx", "MP3",
	}

	m2 := &MusicEntry{
		"3", "Hello world2", "Celion Dion",
		"http://qbox.me/xxxxxxxx", "MP3",
	}
	m3 := &MusicEntry{
		"4", "Hello world3", "Celion Dion",
		"http://qbox.me/xxxxxxxx", "MP3",
	}

	mm.Add(m1)
	mm.Add(m2)
	mm.Add(m3)

	tmp := mm.Remove(0)
	gm0, _ := mm.Get(0)
	gm1, _ := mm.Get(1)
	if gm0.Name != "Hello world2" || gm1.Name != "Hello world3" ||
		tmp.Name != "Hello world1" {
		t.Error("mm.Remove(0) failed")
	}

	mm.Add(tmp)

	tmp = mm.Remove(1)
	if tmp.Name != "Hello world3" || gm0.Name != "Hello world2" ||
		gm1.Name != "Hello world1" {
		t.Error("mm.Remove(1) failed")
		t.Error(gm0)
		t.Error(gm1)
		t.Error(tmp)
	}

	mm.Add(tmp)
	gm0, _ = mm.Get(0)
	gm1, _ = mm.Get(1)
	//gm2, _ := mm.Get(2)

	tmp = mm.RemoveByName("Hello world1")
	gm0, _ = mm.Get(0)
	gm1, _ = mm.Get(1)

	if tmp.Name != "Hello world1" || gm0.Name != "Hello world2" ||
		gm1.Name != "Hello world3" {
		t.Error("mm.Remove(1) failed")
		t.Error(gm0)
		t.Error(gm1)
		t.Error(tmp)
	}
}
