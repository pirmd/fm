package tui

import (
)

// ListWidget represents a scrollable list 
type ListWidget struct {
    // ListWidget is a TextWidget
    *TextWidget
}

// NewListWidget creates new ListWidget, taken its size as parameter
func NewListWidget() *ListWidget {
    return &ListWidget{TextWidget: NewTextWidget()}
}

// SetCursor sets the cursor for viewing
func (l *ListWidget) SetCursor(y int) {
    l.TextWidget.Scroll(0, y)
}

// SetCursorDown moves cursor down by one line
func (l *ListWidget) SetCursorDown() {
    l.TextWidget.Scroll(0, l.vy-1)
}

// SetCursorUp moves cursor up by one line
func (l *ListWidget) SetCursorUp() {
    l.TextWidget.Scroll(0, l.vy+1)
}

// GetCursor gets the line where the viewing cursor is set
func (l *ListWidget) GetCursor() int {
    return l.vy
}
