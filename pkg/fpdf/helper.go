// Package fpdf provides helper utilities for accessing the underlying Fpdf interface
// from Maroto providers for advanced drawing operations.
package fpdf

import (
	"github.com/flanksource/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/flanksource/maroto/v2/pkg/core"
)

// GetFpdf safely extracts the underlying gofpdfwrapper.Fpdf interface from a Provider.
// Returns the Fpdf interface and true if successful, nil and false otherwise.
//
// Example usage:
//   if fpdf, ok := fpdf.GetFpdf(provider); ok {
//       // Use fpdf methods directly
//       fpdf.SetFillColor(255, 0, 0)
//       fpdf.Rect(10, 10, 50, 20, "F")
//   }
func GetFpdf(provider core.Provider) (gofpdfwrapper.Fpdf, bool) {
	if provider == nil {
		return nil, false
	}
	
	fpdfInterface := provider.GetFpdf()
	if fpdfInterface == nil {
		return nil, false
	}
	
	fpdf, ok := fpdfInterface.(gofpdfwrapper.Fpdf)
	return fpdf, ok
}

// GetFpdfFromMaroto safely extracts the underlying gofpdfwrapper.Fpdf interface from a Maroto instance.
// This is a convenience method that combines GetProvider() and GetFpdf().
//
// Example usage:
//   if fpdf, ok := fpdf.GetFpdfFromMaroto(maroto); ok {
//       // Use fpdf methods directly
//       fpdf.SetDrawColor(0, 255, 0)
//       fpdf.Line(0, 0, 100, 100)
//   }
func GetFpdfFromMaroto(maroto core.Maroto) (gofpdfwrapper.Fpdf, bool) {
	if maroto == nil {
		return nil, false
	}
	
	provider := maroto.GetProvider()
	return GetFpdf(provider)
}

// DrawingHelper provides convenience methods for common drawing operations.
type DrawingHelper struct {
	fpdf gofpdfwrapper.Fpdf
}

// NewDrawingHelper creates a new DrawingHelper from a Provider.
// Returns nil if the provider doesn't support Fpdf interface.
func NewDrawingHelper(provider core.Provider) *DrawingHelper {
	if fpdf, ok := GetFpdf(provider); ok {
		return &DrawingHelper{fpdf: fpdf}
	}
	return nil
}

// NewDrawingHelperFromMaroto creates a new DrawingHelper from a Maroto instance.
// Returns nil if the maroto instance doesn't support Fpdf interface.
func NewDrawingHelperFromMaroto(maroto core.Maroto) *DrawingHelper {
	if fpdf, ok := GetFpdfFromMaroto(maroto); ok {
		return &DrawingHelper{fpdf: fpdf}
	}
	return nil
}

// DrawRect draws a rectangle with the specified fill and border.
// styleStr can be:
// - "D" or "" for border only
// - "F" for filled only  
// - "DF" or "FD" for both border and fill
func (dh *DrawingHelper) DrawRect(x, y, w, h float64, styleStr string) {
	if dh.fpdf != nil {
		dh.fpdf.Rect(x, y, w, h, styleStr)
	}
}

// SetFillColor sets the fill color for subsequent drawing operations.
func (dh *DrawingHelper) SetFillColor(r, g, b int) {
	if dh.fpdf != nil {
		dh.fpdf.SetFillColor(r, g, b)
	}
}

// SetDrawColor sets the draw/border color for subsequent drawing operations.
func (dh *DrawingHelper) SetDrawColor(r, g, b int) {
	if dh.fpdf != nil {
		dh.fpdf.SetDrawColor(r, g, b)
	}
}

// DrawLine draws a line from (x1, y1) to (x2, y2).
func (dh *DrawingHelper) DrawLine(x1, y1, x2, y2 float64) {
	if dh.fpdf != nil {
		dh.fpdf.Line(x1, y1, x2, y2)
	}
}

// GetFpdf returns the underlying Fpdf interface for direct access to all methods.
func (dh *DrawingHelper) GetFpdf() gofpdfwrapper.Fpdf {
	return dh.fpdf
}

// DrawPolygon draws a polygon using the provided points.
// points should be in format [][]float64{{x1,y1}, {x2,y2}, ...}
// styleStr can be "D", "F", "DF", or "FD"
func (dh *DrawingHelper) DrawPolygon(points [][]float64, styleStr string) {
	if dh.fpdf == nil || len(points) < 3 {
		return
	}
	
	// Move to first point
	dh.fpdf.MoveTo(points[0][0], points[0][1])
	
	// Draw lines to subsequent points
	for i := 1; i < len(points); i++ {
		dh.fpdf.LineTo(points[i][0], points[i][1])
	}
	
	// Close the polygon
	dh.fpdf.ClosePath()
	
	// Apply the style using DrawPath
	dh.fpdf.DrawPath(styleStr)
}

// DrawCircle draws a circle with the specified style.
// styleStr can be "D", "F", "DF", or "FD"
func (dh *DrawingHelper) DrawCircle(x, y, radius float64, styleStr string) {
	if dh.fpdf != nil {
		dh.fpdf.Circle(x, y, radius, styleStr)
	}
}