package templates

import "embed"

//go:embed *
var TemplateFiles embed.FS

//go:embed static/*
var StaticFiles embed.FS
