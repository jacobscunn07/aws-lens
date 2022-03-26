package ui

type PanelView interface {
	Render() string
}

type PanelData interface {
}

type SidePanelView interface {
	SidePanelViewName() string
	PanelView
}

type SidePanelData interface {
	SidePanelDataName() string
	PanelData
}

type DetailPanelView interface {
	DetailPanelViewName() string
	PanelView
}
