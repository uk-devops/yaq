package pipeline

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Structured data", func() {
	var m1, m2 GenericMap
	var a1, a2 GenericArray
	var sd StructuredData

	BeforeEach(func() {
		m1 = GenericMap{
			"a": 1,
			"b": GenericMap{"c": 2},
		}
		m2 = GenericMap{
			"d": 3,
			"e": GenericMap{"f": 4},
		}
		a1 = GenericArray{
			GenericMap{"a": 1},
			GenericMap{"b": 2},
		}
		a2 = GenericArray{
			GenericMap{"c": 3},
			GenericMap{"d": 4},
		}
	})

	Context("add map to empty initial data", func() {
		BeforeEach(func() {
			sd = StructuredData{}
		})
		It("returns the map", func() {
			sd.Append(m2)
			Expect(sd.Data).To(Equal(m2))
		})
	})

	Context("2 maps with different keys", func() {
		BeforeEach(func() {
			sd = StructuredData{Data: m1}
		})
		It("returns a map with all keys", func() {
			sd.Append(m2)
			Expect(sd.Data).To(Equal(GenericMap{
				"a": 1,
				"b": GenericMap{"c": 2},
				"d": 3,
				"e": GenericMap{"f": 4},
			}))
		})
	})

	Context("2 maps with repeating keys", func() {
		var m3 GenericMap

		BeforeEach(func() {
			sd.Data = nil
			sd = StructuredData{Data: m1}
			m3 = GenericMap{
				"a": 3,
				"b": GenericMap{"c": 4},
			}
		})
		It("returns a map with keys from second map taking priority", func() {
			sd.Append(m3)
			Expect(sd.Data).To(Equal(GenericMap{
				"a": 3,
				"b": GenericMap{"c": 4},
			}))
		})
	})

	Context("add array to empty initial data", func() {
		BeforeEach(func() {
			sd = StructuredData{}
		})
		It("returns the array", func() {
			sd.Append(a2)
			Expect(sd.Data).To(Equal(a2))
		})
	})

	Context("2 arrays", func() {
		BeforeEach(func() {
			sd = StructuredData{Data: a1}
		})
		It("returns the concatenated array", func() {
			sd.Append(a2)
			Expect(sd.Data).To(Equal(append(a1, a2...)))
		})
	})

	Context("1 map 1 array", func() {
		var a GenericArray

		BeforeEach(func() {
			sd = StructuredData{Data: m1}
			sd.Append(a2)
		})
		It("returns an array", func() {
			var ok bool
			a, ok = sd.Data.(GenericArray)
			Expect(ok).To(BeTrue())
		})
		It("the first element is the map", func() {
			Expect(a[0]).To(Equal(m1))
		})
		It("the second and third elements are the array elements", func() {
			Expect(a[1]).To(Equal(a2[0]))
			Expect(a[2]).To(Equal(a2[1]))
		})
	})

	Context("1 array 1 map", func() {
		var a GenericArray

		BeforeEach(func() {
			sd = StructuredData{Data: a1}
			sd.Append(m2)
		})
		It("returns an array", func() {
			var ok bool
			a, ok = sd.Data.(GenericArray)
			Expect(ok).To(BeTrue())
		})
		It("the first and second elements are the array elements", func() {
			Expect(a[0]).To(Equal(a1[0]))
			Expect(a[1]).To(Equal(a1[1]))
		})
		It("the third element is the map", func() {
			Expect(a[2]).To(Equal(m2))
		})
	})

})
