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

package pkg_test

import (
	. "github.com/mudler/luet/pkg/package"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Database", func() {

	db := NewInMemoryDatabase(false)
	Context("Simple package", func() {
		a := NewPackage("A", ">=1.0", []*DefaultPackage{}, []*DefaultPackage{})
		//	a1 := NewPackage("A", "1.0", []*DefaultPackage{}, []*DefaultPackage{})
		//		a11 := NewPackage("A", "1.1", []*DefaultPackage{}, []*DefaultPackage{})
		//		a01 := NewPackage("A", "0.1", []*DefaultPackage{}, []*DefaultPackage{})
		It("Saves and get data back correctly", func() {

			ID, err := db.CreatePackage(a)
			Expect(err).ToNot(HaveOccurred())

			pack, err := db.GetPackage(ID)
			Expect(err).ToNot(HaveOccurred())

			Expect(pack).To(Equal(a))

		})

		It("Gets all", func() {

			ids := db.GetPackages()

			Expect(ids).To(Equal([]string{"A-->=1.0"}))

		})
		It("Find packages", func() {

			pack, err := db.FindPackage(a)
			Expect(err).ToNot(HaveOccurred())
			Expect(pack).To(Equal(a))

		})

		It("Find best package candidate", func() {
			db := NewInMemoryDatabase(false)
			a := NewPackage("A", "1.0", []*DefaultPackage{}, []*DefaultPackage{})
			a1 := NewPackage("A", "1.1", []*DefaultPackage{}, []*DefaultPackage{})
			a3 := NewPackage("A", "1.3", []*DefaultPackage{}, []*DefaultPackage{})
			_, err := db.CreatePackage(a)
			Expect(err).ToNot(HaveOccurred())

			_, err = db.CreatePackage(a1)
			Expect(err).ToNot(HaveOccurred())

			_, err = db.CreatePackage(a3)
			Expect(err).ToNot(HaveOccurred())
			s := NewPackage("A", ">=1.0", []*DefaultPackage{}, []*DefaultPackage{})

			pack, err := db.FindPackageCandidate(s)
			Expect(err).ToNot(HaveOccurred())
			Expect(pack).To(Equal(a3))

		})
	})

})
