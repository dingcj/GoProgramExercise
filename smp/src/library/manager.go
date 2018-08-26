package library

import (
	"errors"
)

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range")
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}

	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removeMusic := m.musics[index]
	if index > 0 && index < len(m.musics)-1 {
		m.musics = append(m.musics[:index], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = m.musics[1:]
	} else {
		m.musics = m.musics[:index-1]
	}

	return &removeMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {

	for idx, music := range m.musics {
		if music.Name == name {
			m.Remove(idx)
			return &music
		}
	}

	return nil
}
