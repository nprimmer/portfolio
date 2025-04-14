// Package memory provides a simplified simulation of memory paging.
package main

import (
	"os"
	"sync"
	"time"
)

const (
	// PageSize defines the size of each page in bytes.
	PageSize = 1024
	// MemoryLimit defines the maximum number of pages that can be held in memory at one time.
	MemoryLimit = 4
	// TotalPages defines the total number of pages in the simulated file.
	TotalPages = 1000
)

// Page represents a single page of memory.
type Page struct {
	ID   int
	Data []byte
}

// Memory manages a fixed-size set of pages, loading them from a file as needed.
type Memory struct {
	mu       sync.Mutex
	pages    map[int]*Page
	order    []int // Order of page IDs for replacement policy
	capacity int
	file     *os.File
}

// NewMemory creates a new Memory instance with a given file for backing storage.
func NewMemory(file *os.File) *Memory {
	return &Memory{
		pages:    make(map[int]*Page),
		order:    []int{},
		capacity: MemoryLimit,
		file:     file,
	}
}

// loadPage loads a page from the file into memory, evicting another page if necessary.
func (m *Memory) loadPage(pageID int) (*Page, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if the page is already in memory
	if page, exists := m.pages[pageID]; exists {
		m.moveToEnd(pageID)
		return page, nil
	}

	// Evict a page if we are at capacity
	if len(m.pages) >= m.capacity {
		m.evictPage()
	}

	// Load the page from the file
	page := &Page{
		ID:   pageID,
		Data: make([]byte, PageSize),
	}
	offset := int64(pageID) * PageSize
	_, err := m.file.ReadAt(page.Data, offset)
	if err != nil {
		return nil, err
	}

	// Simulate disk access delay
	time.Sleep(100 * time.Millisecond)

	// Add the page to memory and update the order
	m.pages[pageID] = page
	m.order = append(m.order, pageID)
	return page, nil
}

// moveToEnd moves a page ID to the end of the order slice to simulate LRU policy.
func (m *Memory) moveToEnd(pageID int) {
	for i, id := range m.order {
		if id == pageID {
			m.order = append(m.order[:i], m.order[i+1:]...)
			break
		}
	}
	m.order = append(m.order, pageID)
}

// evictPage evicts the least recently used page from memory.
func (m *Memory) evictPage() {
	if len(m.order) == 0 {
		return
	}

	evictID := m.order[0]
	m.order = m.order[1:]
	delete(m.pages, evictID)
}

// ReadAddress reads a byte from the specified address.
func (m *Memory) ReadAddress(address int) (byte, error) {
	pageID := address / PageSize
	offset := address % PageSize
	page, err := m.loadPage(pageID)
	if err != nil {
		return 0, err
	}

	return page.Data[offset], nil
}

// ReadPage reads an entire page by its ID.
func (m *Memory) ReadPage(pageID int) ([]byte, error) {
	page, err := m.loadPage(pageID)
	if err != nil {
		return nil, err
	}
	return page.Data, nil
}
