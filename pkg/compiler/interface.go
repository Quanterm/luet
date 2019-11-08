// Copyright © 2019 Ettore Di Giacinto <mudler@gentoo.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package compiler

import pkg "github.com/mudler/luet/pkg/package"

type Compiler interface {
	Compile(CompilationSpec) (*Artifact, error)
	FromPackage(pkg.Package) (CompilationSpec, error)

	SetBackend(CompilerBackend)
	GetBackend() CompilerBackend
}

type CompilerBackendOptions struct {
	ImageName      string
	SourcePath     string
	DockerFileName string
	Destination    string
}

type CompilerBackend interface {
	BuildImage(CompilerBackendOptions) error
	ExportImage(CompilerBackendOptions) error
	RemoveImage(CompilerBackendOptions) error
	ImageDefinitionToTar(CompilerBackendOptions) error
}

// CompilationSpec represent a compilation specification derived from a package
type CompilationSpec interface {
	RenderBuildImage() (string, error)
	WriteBuildImageDefinition(string) error

	RenderStepImage(image string) (string, error)
	WriteStepImageDefinition(fromimage, path string) error

	GetPackage() pkg.Package
	BuildSteps() []string

	GetSeedImage() string
	SetSeedImage(string)

	GetImage() string
	SetImage(string)

	SetOutputPath(string)
	GetOutputPath() string
	Rel(string) string

	GetPreBuildSteps() []string
}
