package main

import (
	"bytes"
	"github.com/griffin702/service/tools"
	"runtime"
	"text/template"
	"time"
)

var (
	Version   = "v1.0.8"
	BuildTime = tools.Tools.TimeFormat(time.Now(), "Y-m-d H:i:s")
)

// VersionOptions include version
type VersionOptions struct {
	GitCommit string
	Version   string
	BuildTime string
	GoVersion string
	Os        string
	Arch      string
}

var versionTemplate = `Version: {{.Version}} | Go version: {{.GoVersion}} | BuildTime: {{.BuildTime}} | OS/Arch: {{.Os}}/{{.Arch}}`

func getVersion() string {
	var doc bytes.Buffer
	vo := VersionOptions{
		Version:   Version,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
	tmpl, _ := template.New("version").Parse(versionTemplate)
	tmpl.Execute(&doc, vo)
	return doc.String()
}
